terraform {
  required_providers {
    librenms = {
      version = "0.0.3"
      source  = "github.com/rukas/librenms"
    }
  }
}

provider "librenms" {}

resource "librenms_device" "localhost" {
  hostname         = "localhost"
  community_string = "community"
  snmp_disable     = 0
  snmp_port        = 161
  snmp_version     = "v2c"
}
