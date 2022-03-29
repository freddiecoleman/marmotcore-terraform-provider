package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNodeMarmotCore(t *testing.T) {

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNodeMarmotCore,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"marmotcore_node.foo", "sample_attribute", regexp.MustCompile("^ba")),
				),
			},
		},
	})
}

const testAccNodeMarmotCore = `
resource "marmotcore_node" "foo" {
  sample_attribute = "bar"
}
`
