package provider

import (
	"regexp"
	"testing"

	marmotcoreclient "github.com/freddiecoleman/marmotcore-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

type MockClient struct {
	Protocol   string
	Host       string
	Port       string
	ApiVersion string
}

var CreateNodeFunc func(createNode *marmotcoreclient.CreateNode) (marmotcoreclient.CreateNodeResponse, error)

func (*MockClient) CreateNode(createNode *marmotcoreclient.CreateNode) (marmotcoreclient.CreateNodeResponse, error) {
	return CreateNodeFunc(createNode)
}

func TestAccNodeMarmotCore(t *testing.T) {
	CreateNodeFunc = func(createNode *marmotcoreclient.CreateNode) (marmotcoreclient.CreateNodeResponse, error) {
		return marmotcoreclient.CreateNodeResponse{
			NodeId: "amazing-123",
		}, nil
	}

	GetClient = func(protocol string, host string, port string, apiVersion string) (Client, error) {
		return &MockClient{
			Protocol:   "http",
			Host:       "localhost",
			Port:       "3000",
			ApiVersion: "v1",
		}, nil
	}

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNodeMarmotCore,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"marmotcore_node.foo", "id", regexp.MustCompile("^amazing-123")),
					resource.TestMatchResourceAttr(
						"marmotcore_node.foo", "region", regexp.MustCompile("^us-west-2")),
					resource.TestMatchResourceAttr(
						"marmotcore_node.foo", "instance_type", regexp.MustCompile("^node.small")),
					resource.TestMatchResourceAttr(
						"marmotcore_node.foo", "chia_version", regexp.MustCompile("^1.3.*")),
					resource.TestMatchResourceAttr(
						"marmotcore_node.foo", "network", regexp.MustCompile("^testnet")),
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
