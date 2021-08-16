package librenms

import (
	"context"
	"time"

	lnms "github.com/rukas/librenms-go-client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDevice() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDeviceCreate,
		ReadContext:   resourceDeviceRead,
		UpdateContext: resourceDeviceUpdate,
		DeleteContext: resourceDeviceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				Optional: false,
				ForceNew: true,
			},
			"snmp_port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: false,
				Optional: true,
				Default:  161,
			},
			"snmp_version": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Default:  "v2c",
			},
			"snmp_disable": &schema.Schema{
				Type:     schema.TypeInt,
				Required: false,
				Optional: true,
				Default:  0,
			},
			"community_string": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*lnms.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	ds := lnms.Device{}
	ds.Hostname = d.Get("hostname").(string)
	ds.Port = d.Get("snmp_port").(int)
	ds.Version = d.Get("snmp_version").(string)
	ds.SnmpDisable = d.Get("snmp_disable").(int)
	ds.Community = d.Get("community_string").(string)

	o, err := c.CreateDevice(ds)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(o.Hostname)

	resourceDeviceRead(ctx, d, m)

	return diags
}

func resourceDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*lnms.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	deviceID := d.Id()

	device, err := c.GetDevice(deviceID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(device.Hostname)
	d.Set("hostname", device.Hostname)
	d.Set("snmp_port", device.Port)
	d.Set("snmp_version", device.Snmpver)
	d.Set("snmp_disable", device.SnmpDisable)
	d.Set("community_string", device.Community)

	return diags
}

func resourceDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*lnms.Client)

	if d.HasChange("snmp_port") || d.HasChange("community_string") {

		ds := lnms.Device{}
		ds.Hostname = d.Get("hostname").(string)
		ds.Port = d.Get("snmp_port").(int)
		ds.Version = d.Get("snmp_version").(string)
		ds.SnmpDisable = d.Get("snmp_disable").(int)
		ds.Community = d.Get("community_string").(string)

		_, err := c.UpdateDevice(ds)
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceDeviceRead(ctx, d, m)
}

func resourceDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*lnms.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	hostname := d.Get("hostname").(string)

	err := c.DeleteDevice(hostname)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
