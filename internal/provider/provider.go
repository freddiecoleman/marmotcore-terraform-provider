package provider

import (
	"context"

	marmotcoreclient "github.com/freddiecoleman/marmotcore-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"protocol": {
					Type:     schema.TypeString,
					Required: true,
				},
				"host": {
					Type:     schema.TypeString,
					Required: true,
				},
				"port": {
					Type:     schema.TypeString,
					Required: true,
				},
				"api_version": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"scaffolding_data_source": dataSourceScaffolding(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"marmotcore_node": nodeMarmotCore(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func GetClient(protocol string, host string, port string, apiVersion string) (marmotcoreclient.MarmotcoreClient, error) {
	return marmotcoreclient.MarmotcoreClient{
		Protocol:   protocol,
		Host:       host,
		Port:       port,
		ApiVersion: apiVersion,
	}, nil
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		protocol := d.Get("protocol").(string)
		host := d.Get("host").(string)
		port := d.Get("port").(string)
		apiVersion := d.Get("api_version").(string)

		c, err := GetClient(protocol, host, port, apiVersion)

		if err != nil {
			diag.Errorf("Failed to create Marmotcore client")
		}

		return c, nil
	}
}
