package anchore

import (
	"context"

	"github.com/excellaco/anchore-client-go/client"
	"github.com/excellaco/anchore-client-go/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAnchoreRegistry() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAnchoreRegistryCreate,
		ReadContext:   resourceAnchoreRegistryRead,
		UpdateContext: resourceAnchoreRegistryUpdate,
		DeleteContext: resourceAnchoreRegistryDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"url": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAnchoreRegistryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	anchoreClient := meta.(client.AnchoreClient)

	url := d.Get("url").(string)
	name := d.Get("name").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	typeStr := d.Get("type").(string)

	registry, err := anchoreClient.RegistryCreate(types.Registry{
		URL:      url,
		Name:     name,
		Username: username,
		Password: password,
		Type:     typeStr,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(registry[0].URL)

	return resourceAnchoreRegistryRead(ctx, d, meta)
}

func resourceAnchoreRegistryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	anchoreClient := meta.(client.AnchoreClient)
	url := d.Id()

	registry, err := anchoreClient.RegistryRead(&url)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(registry.URL)

	if err := d.Set("url", registry.URL); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", registry.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("username", registry.Username); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("type", registry.Type); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceAnchoreRegistryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	anchoreClient := meta.(client.AnchoreClient)
	url := d.Id()

	registry, err := anchoreClient.RegistryRead(&url)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = anchoreClient.RegistryUpdate(types.Registry{
		URL:      registry.URL,
		Name:     d.Get("name").(string),
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		Type:     d.Get("type").(string),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceAnchoreRegistryRead(ctx, d, meta)
}

func resourceAnchoreRegistryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	anchoreClient := meta.(client.AnchoreClient)
	url := d.Id()

	err := anchoreClient.RegistryDelete(&url)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
