package acc_tests

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceMetaport(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      validateResourceDestroyed("metaport", "v1/metaports"),
		Steps: []resource.TestStep{
			{
				Config: testAccMetaportStep1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"pfptmeta_metaport.metaport", "id", regexp.MustCompile("^mp-[\\d]+$"),
					),
					resource.TestCheckResourceAttr(
						"pfptmeta_metaport.metaport", "name", "metaport name",
					),
					resource.TestCheckResourceAttr(
						"pfptmeta_metaport.metaport", "description", "metaport description",
					),
					resource.TestCheckTypeSetElemAttrPair(
						"pfptmeta_metaport.metaport", "mapped_elements.*", "pfptmeta_network_element.mapped-subnet", "id",
					),
					resource.TestCheckResourceAttr(
						"pfptmeta_metaport.metaport", "allow_support", "true",
					),
					resource.TestCheckTypeSetElemAttrPair(
						"pfptmeta_metaport.metaport", "notification_channels.*", "pfptmeta_notification_channel.mail", "id",
					),
				),
			},
			{
				Config: testAccMetaportStep2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"pfptmeta_metaport.metaport", "name", "metaport name1",
					),
					resource.TestCheckResourceAttr(
						"pfptmeta_metaport.metaport", "description", "metaport description1",
					),
					resource.TestCheckResourceAttr("pfptmeta_metaport.metaport", "mapped_elements.#", "0"),
					resource.TestCheckResourceAttr("pfptmeta_metaport.metaport", "notification_channels.#", "0"),
					resource.TestCheckResourceAttr("pfptmeta_metaport.metaport", "allow_support", "false"),
				),
			},
		},
	})
}

func TestAccDataSourceMetaport(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      validateResourceDestroyed("metaport", "v1/metaports"),
		Steps: []resource.TestStep{
			{
				Config: testAccMetaportStep1 + testAccMetaportByIDDataSource,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"data.pfptmeta_metaport.metaport", "id", regexp.MustCompile("^mp-[\\d]+$"),
					),
					resource.TestCheckResourceAttr(
						"data.pfptmeta_metaport.metaport", "name", "metaport name",
					),
					resource.TestCheckResourceAttr(
						"data.pfptmeta_metaport.metaport", "description", "metaport description",
					),
					resource.TestCheckTypeSetElemAttrPair(
						"data.pfptmeta_metaport.metaport", "mapped_elements.*", "pfptmeta_network_element.mapped-subnet", "id",
					),
					resource.TestCheckResourceAttr(
						"data.pfptmeta_metaport.metaport", "allow_support", "true",
					),
				),
			},
			{
				Config: testAccMetaportStep1 + testAccMetaportByNameDataSource,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"data.pfptmeta_metaport.metaport", "id", regexp.MustCompile("^mp-[\\d]+$"),
					),
					resource.TestCheckResourceAttr(
						"data.pfptmeta_metaport.metaport", "name", "metaport name",
					),
					resource.TestCheckResourceAttr(
						"data.pfptmeta_metaport.metaport", "description", "metaport description",
					),
					resource.TestCheckTypeSetElemAttrPair(
						"data.pfptmeta_metaport.metaport", "mapped_elements.*", "pfptmeta_network_element.mapped-subnet", "id",
					),
					resource.TestCheckResourceAttr(
						"data.pfptmeta_metaport.metaport", "allow_support", "true",
					),
				),
			},
		},
	})
}

const testAccMetaportStep1 = `
resource "pfptmeta_network_element" "mapped-subnet" {
  name           = "ms"
  mapped_subnets = ["0.0.0.0/0"]
}

resource "pfptmeta_notification_channel" "mail" {
  name        = "mail-channel"
  description = "mail channel description"
  email_config {
    recipients = ["user1@example.com", "user2@example.com"]
  }
}

resource "pfptmeta_metaport" "metaport" {
  name                  = "metaport name"
  description           = "metaport description"
  mapped_elements       = [pfptmeta_network_element.mapped-subnet.id]
  notification_channels = [pfptmeta_notification_channel.mail.id]
}
`

const testAccMetaportStep2 = `
resource "pfptmeta_network_element" "mapped-subnet" {
  name           = "ms"
  mapped_subnets = ["0.0.0.0/0"]
}

resource "pfptmeta_metaport" "metaport" {
  name                  = "metaport name1"
  description           = "metaport description1"
  allow_support         = false
  mapped_elements       = []
}
`

const testAccMetaportByIDDataSource = `

data "pfptmeta_metaport" "metaport" {
  id = pfptmeta_metaport.metaport.id
}`

const testAccMetaportByNameDataSource = `

data "pfptmeta_metaport" "metaport" {
  name = "metaport name"
}`
