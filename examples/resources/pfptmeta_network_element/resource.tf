resource "pfptmeta_network_element" "mapped-subnet" {
  name           = "mapped subnet name"
  description    = "some details about the mapped subnet"
  mapped_subnets = ["0.0.0.0/0", "10.20.30.0/24"]
  tags = {
    tag_name1 = "tag_value1"
    tag_name2 = "tag_value2"
  }
  mapped_domains {
    name          = "test.com"
    mapped_domain = "test.com"
  }
  mapped_domains {
    name          = "test1.com"
    mapped_domain = "test1.com"
  }
  mapped_hosts {
    name        = "host.com"
    mapped_host = "10.0.0.1"
  }
  mapped_hosts {
    name        = "host1.com"
    mapped_host = "10.0.0.2"
  }
}

resource "pfptmeta_network_element" "mapped-service" {
  name           = "mapped service name"
  description    = "some details about the mapped service"
  mapped_service = "mapped.service.com"
  tags = {
    tag_name1 = "tag_value1"
    tag_name2 = "tag_value2"
  }
  aliases = ["alias1.mapped.service.com", "alias2.mapped.service.com"]
}

resource "pfptmeta_network_element" "device" {
  name        = "my device"
  description = "some details about the device"
  owner_id    = "usr-12345"
  platform    = "Linux"
  aliases     = ["alias1.device.com", "alias2.device.com"]
}