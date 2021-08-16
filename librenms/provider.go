package librenms

import (
	"context"

	lnms "github.com/rukas/librenms-go-client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"librenms_host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("LIBRENMS_HOST", nil),
			},
			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("LIBRENMS_API_KEY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"librenms_device": resourceDevice(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			// "librenms_device": dataSourceDevice(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	host := d.Get("librenms_host").(string)
	apiKey := d.Get("api_key").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if apiKey != "" {
		c, err := lnms.NewClient(&host, &apiKey)
		if err != nil {
			return nil, diag.FromErr(err)
		}
		return c, diags
	}

	c, err := lnms.NewClient(nil, nil)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}
