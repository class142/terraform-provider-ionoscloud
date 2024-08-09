package s3

import (
	"context"
	"fmt"
	"net/http"

	s3 "github.com/ionos-cloud/sdk-go-s3"
)

// DeleteRequest represents a request to delete an object from an S3 general purpose bucket.
type DeleteRequest struct {
	Bucket              string
	Key                 string
	VersionID           string
	ForceDestroy        bool
	IgnoreObjectsErrors bool
}

// DeleteAllObjectVersions deletes all versions of a specified key from an S3 general purpose bucket.
// If key is empty then all versions of all objects are deleted.
// Set `force` to `true` to override any S3 object lock protections on object lock enabled buckets.
// Returns the number of objects deleted.
// Use `emptyBucket` to delete all versions of all objects in a bucket.
func DeleteAllObjectVersions(ctx context.Context, client *s3.APIClient, req *DeleteRequest) (int, error) {
	var (
		objCount int
		lastErr  error
	)

	pages := NewListObjectVersionsPaginator(client, &ListObjectVersionsInput{
		Bucket: req.Bucket,
		Prefix: req.Key,
	})

	for pages.HasMorePages() {
		page, err := pages.NextPage(ctx)
		if err != nil {
			return objCount, err
		}

		var count int
		count, lastErr = deleteVersionsPage(ctx, client, page.Versions, req.Bucket, req.Key, req.ForceDestroy)
		objCount += count
	}

	if lastErr != nil {
		if !req.IgnoreObjectsErrors {
			return objCount, lastErr
		}
		lastErr = nil
	}

	pages = NewListObjectVersionsPaginator(client, &ListObjectVersionsInput{
		Bucket: req.Bucket,
		Prefix: req.Key,
	})

	for pages.HasMorePages() {
		page, err := pages.NextPage(ctx)
		if err != nil {
			return objCount, err
		}

		var count int
		count, lastErr = deleteMarkersPage(ctx, client, page.DeleteMarkers, req.Bucket, req.Key)
		objCount += count
	}

	if lastErr != nil {
		if !req.IgnoreObjectsErrors {
			return objCount, fmt.Errorf("deleting at least one S3 Object delete marker, last error: %w", lastErr)
		}
	}

	return objCount, nil
}

func deleteObject(ctx context.Context, client *s3.APIClient, req *DeleteRequest) (*s3.APIResponse, error) {
	r := client.ObjectsApi.DeleteObject(ctx, req.Bucket, req.Key)
	if req.VersionID != "" {
		r = r.VersionId(req.VersionID)
	}

	if req.ForceDestroy {
		r = r.XAmzBypassGovernanceRetention(true)
	}

	_, apiResponse, err := r.Execute()
	return apiResponse, err
}

func deleteVersionsPage(ctx context.Context, client *s3.APIClient, versions *[]s3.ObjectVersion, bucket, key string, force bool) (int, error) {
	var (
		objCount int
		lastErr  error
	)

	if versions == nil {
		return objCount, nil
	}

	for _, v := range *versions {
		if key != toString(v.Key) {
			continue
		}

		apiResponse, err := deleteObject(ctx, client, &DeleteRequest{
			Bucket:       bucket,
			Key:          key,
			VersionID:    toString(v.VersionId),
			ForceDestroy: force,
		})

		if err == nil {
			objCount++
			continue
		}

		// If the object is locked by Object Lock, we need to disable the legal hold before deleting it
		if httpForbidden(apiResponse) && force {
			success, err := tryDisableLegalHold(ctx, client, bucket, key, toString(v.VersionId))
			if err != nil {
				lastErr = err
				continue
			}

			if !success {
				lastErr = err
				continue
			}

			if _, err = deleteObject(ctx, client, &DeleteRequest{
				Bucket:       bucket,
				Key:          key,
				VersionID:    toString(v.VersionId),
				ForceDestroy: force,
			}); err != nil {
				lastErr = err
				continue
			}

			objCount++
		}
		lastErr = err
	}

	return objCount, lastErr
}

func deleteMarkersPage(ctx context.Context, client *s3.APIClient, markers *[]s3.DeleteMarkerEntry, bucket, key string) (int, error) {
	var (
		objCount int
		lastErr  error
	)

	if markers == nil {
		return objCount, nil
	}

	for _, m := range *markers {
		if key != toString(m.Key) {
			continue
		}

		_, err := deleteObject(ctx, client, &DeleteRequest{
			Bucket:    bucket,
			Key:       key,
			VersionID: *m.VersionId,
		})

		if err == nil {
			objCount++
			continue
		}

		lastErr = err
	}

	return objCount, lastErr

}

func tryDisableLegalHold(ctx context.Context, client *s3.APIClient, bucket, key, versionID string) (bool, error) {
	output, _, err := client.ObjectLockApi.GetObjectLegalHold(ctx, bucket, key).VersionId(versionID).Execute()
	if err != nil {
		return false, err
	}

	if *output.Status == "OFF" {
		return false, nil
	}

	_, err = client.ObjectLockApi.PutObjectLegalHold(ctx, bucket, key).VersionId(versionID).
		ObjectLegalHoldConfiguration(s3.ObjectLegalHoldConfiguration{
			Status: strPtr("OFF"),
		}).Execute()

	if err != nil {
		return false, err
	}

	return true, nil
}

func httpForbidden(response *s3.APIResponse) bool {
	return response != nil && response.StatusCode == http.StatusForbidden
}

func strPtr(s string) *string {
	return &s
}