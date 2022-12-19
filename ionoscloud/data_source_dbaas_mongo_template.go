package ionoscloud

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	mongo "github.com/ionos-cloud/sdk-go-dbaas-mongo"
	dbaasService "github.com/ionos-cloud/terraform-provider-ionoscloud/v6/services/dbaas"
)

func dataSourceDbassMongoTemplate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDbassMongoTemplateRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:         schema.TypeString,
				Description:  "The unique ID of the template.",
				ValidateFunc: validation.IsUUID,
				Optional:     true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the template.",
				Optional:    true,
			},
			"partial_match": {
				Type:        schema.TypeBool,
				Description: "Whether partial matching is allowed or not when using the name filter.",
				Default:     false,
				Optional:    true,
			},
			"edition": {
				Type:        schema.TypeString,
				Description: "The edition of the template (e.g. enterprise).",
				Computed:    true,
			},
			"cores": {
				Type:        schema.TypeInt,
				Description: "The number of CPU cores.",
				Computed:    true,
			},
			"ram": {
				Type:        schema.TypeInt,
				Description: "The amount of memory in GB.",
				Computed:    true,
			},
			"storage_size": {
				Type:        schema.TypeInt,
				Description: "The amount of storage size in GB.",
				Computed:    true,
			},
		},
		Timeouts: &resourceDefaultTimeouts,
	}
}

func dataSourceDbassMongoTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(SdkBundle).MongoClient
	id, idOk := d.GetOk("id")
	name, nameOk := d.GetOk("name")

	// Initial checks.
	if idOk && nameOk {
		return diag.FromErr(fmt.Errorf("name and ID cannot be both specified at the same time"))
	}
	if !idOk && !nameOk {
		return diag.FromErr(fmt.Errorf("please provide a template ID or name"))
	}

	retrievedTemplates, _, err := client.GetTemplates(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("an error occurred while fetching dbaas mongo templates: %w", err))
	}

	var templates []mongo.TemplateResponse
	partialMatch := d.Get("partial_match").(bool)
	if retrievedTemplates.Items != nil {
		for _, retrievedTemplate := range *retrievedTemplates.Items {
			// Filter using the template ID or name.
			if (idOk && matchesId(retrievedTemplate, id.(string))) || (nameOk && matchesName(retrievedTemplate, name.(string), partialMatch)) {
				templates = append(templates, retrievedTemplate)
			}
		}
	}
	if templates == nil {
		return diag.FromErr(fmt.Errorf("no DBaaS Mongo Template found with the specified criteria"))
	} else if len(templates) > 1 {
		return diag.FromErr(fmt.Errorf("more than one DBaaS Mongo Template found for the specified search criteria"))
	}
	if err := dbaasService.SetMongoDBTemplateData(d, templates[0]); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

// Check if a template has a specific ID.
func matchesId(template mongo.TemplateResponse, templateId string) bool {
	if template.Id != nil && *template.Id == templateId {
		return true
	}
	return false
}

// Check if a template has a specific name, also supports partial matching.
func matchesName(template mongo.TemplateResponse, name string, partialMatch bool) bool {
	if template.Name != nil {
		if partialMatch {
			if strings.Contains(*template.Name, name) {
				return true
			}
		} else {
			if *template.Name == name {
				return true
			}
		}
	}
	return false
}