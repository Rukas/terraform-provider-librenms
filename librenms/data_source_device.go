package librenms

import (
	"context"
	"fmt"

	lnms "github.com/rukas/librenms-go-client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDevice() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDeviceRead,
		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"snmp_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"snmp_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"snmp_disable": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"community_string": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*lnms.Client)
	fmt.Println("[DEBUG] inside dataSourceDeviceRead")

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	hostname := d.Get("hostname").(string)
	// fmt.Println(hostname)
	device, err := c.GetDevice(hostname)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(device.Hostname)
	d.Set("hostname", device.Hostname)
	d.Set("snmp_port", device.Port)
	d.Set("snmp_version", device.Snmpver)
	d.Set("snmp_disable", device.SnmpDisable)
	d.Set("community_string", device.Community)

	// if err := d.Set("device", device); err != nil {
	// 	return diag.FromErr(err)
	// }

	// d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
