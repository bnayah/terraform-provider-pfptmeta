package acc_tests

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"regexp"
	"testing"
)

func TestAccResourceProtocolGroup(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      validateResourceDestroyed("protocol_group", "v1/protocol_groups"),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProtocolGroup,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"pfptmeta_protocol_group.new_protocol", "id", regexp.MustCompile("^pg-.*$"),
					),
					resource.TestCheckResourceAttr(
						"pfptmeta_protocol_group.new_protocol", "name", "NEW_PROTOCOL",
					),
					resource.TestCheckResourceAttr(
						"pfptmeta_protocol_group.new_protocol", "protocols.0.from_port", "445",
					),
					resource.TestCheckResourceAttr(
						"pfptmeta_protocol_group.new_protocol", "protocols.0.to_port", "445",
					),
					resource.TestCheckResourceAttr(
						"pfptmeta_protocol_group.new_protocol", "protocols.0.proto", "udp",
					),
					resource.TestCheckResourceAttr(
						"pfptmeta_protocol_group.new_protocol", "protocols.1.from_port", "446",
					),
					resource.TestCheckResourceAttr(
						"pfptmeta_protocol_group.new_protocol", "protocols.1.to_port", "446",
					),
					resource.TestCheckResourceAttr(
						"pfptmeta_protocol_group.new_protocol", "protocols.1.proto", "tcp",
					),
				),
			},
		},
	})
}

const testAccResourceProtocolGroup = `
resource "pfptmeta_protocol_group" "new_protocol" {
  name = "NEW_PROTOCOL"
  protocols {
    from_port = 445
    to_port   = 445
    proto     = "udp"
  }
  protocols {
    from_port = 446
    to_port   = 446
    proto     = "tcp"
  }
}
`

func TestAccDataSourceProtocolGroup(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceProtocolGroup,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.pfptmeta_protocol_group.HTTP", "id", "pg-NKMzUnJzalxWZKe",
					),
					resource.TestCheckResourceAttr(
						"data.pfptmeta_protocol_group.HTTP", "name", "HTTP",
					),
					resource.TestCheckResourceAttr(
						"data.pfptmeta_protocol_group.HTTP", "protocols.0.from_port", "80",
					),
					resource.TestCheckResourceAttr(
						"data.pfptmeta_protocol_group.HTTP", "protocols.0.to_port", "80",
					),
					resource.TestCheckResourceAttr(
						"data.pfptmeta_protocol_group.HTTP", "protocols.0.proto", "tcp",
					),
					resource.TestCheckResourceAttr(
						"data.pfptmeta_protocol_group.HTTPS", "id", "pg-D6vrsnKbj7q5Kxl",
					),
					resource.TestCheckResourceAttr(
						"data.pfptmeta_protocol_group.HTTPS", "name", "HTTPS",
					),
					resource.TestCheckResourceAttr(
						"data.pfptmeta_protocol_group.HTTPS", "protocols.0.from_port", "443",
					),
					resource.TestCheckResourceAttr(
						"data.pfptmeta_protocol_group.HTTPS", "protocols.0.to_port", "443",
					),
					resource.TestCheckResourceAttr(
						"data.pfptmeta_protocol_group.HTTPS", "protocols.0.proto", "tcp",
					),
				),
			},
		},
	})
}

const testAccDataSourceProtocolGroup = `
data "pfptmeta_protocol_group" "HTTP" {
  id = "pg-NKMzUnJzalxWZKe"
}

data "pfptmeta_protocol_group" "HTTPS" {
  name = "HTTPS"
}
`
