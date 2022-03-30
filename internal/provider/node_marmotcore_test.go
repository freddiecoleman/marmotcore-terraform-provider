package provider

import (
	"regexp"
	"testing"

	marmotcoreclient "github.com/freddiecoleman/marmotcore-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

type MockClient struct {
	CreateNodeFunc func(createNode *marmotcoreclient.CreateNode) (marmotcoreclient.CreateNodeResponse, error)
}

var CreateNodeFunc func(createNode *marmotcoreclient.CreateNode) (marmotcoreclient.CreateNodeResponse, error)

func (*MockClient) CreateNode(createNode *marmotcoreclient.CreateNode) (marmotcoreclient.CreateNodeResponse, error) {
	return CreateNodeFunc(createNode)
}

func TestAccNodeMarmotCore(t *testing.T) {

	// Todo: overwrite GetClient to inject a mock client which asserts the calls made and responds with a http 200 response including the correct data
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNodeMarmotCore,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"marmotcore_node.foo", "protocol", regexp.MustCompile("^http")),
				),
			},
		},
	})
}

const testAccNodeMarmotCore = `
provider "marmotcore" {
	protocol    = "http"
	host        = "localhost"
	port        = "3000"
	api_version = "v1"
}

resource "marmotcore_node" "foo" {
	region        = "us-west-2"
	instance_type = "node.small"
	chia_version  = "1.3.*"
	network       = "testnet"
}
`
