package provider

import (
	"context"

	marmotcoreclient "github.com/freddiecoleman/marmotcore-client"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func nodeMarmotCore() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Marmotcore Cloud Full Node",

		CreateContext: nodeMarmotCoreCreate,
		ReadContext:   nodeMarmotCoreRead,
		UpdateContext: nodeMarmotCoreUpdate,
		DeleteContext: nodeMarmotCoreDelete,

		Schema: map[string]*schema.Schema{
			"region": {
				Description: "Region in which to deploy marmotcore node.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"instance_type": {
				Description: "Size of marmotcore node to deploy",
				Type:        schema.TypeString,
				Required:    true,
			},
			"chia_version": {
				Description: "Version of Chia to run on marmotcore node",
				Type:        schema.TypeString,
				Required:    true,
			},
			"network": {
				Description: "Network to run marmotcore node against. Typically mainnet or testnet.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func nodeMarmotCoreCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Client)

	var diags diag.Diagnostics

	result, err := client.CreateNode(&marmotcoreclient.CreateNode{
		Region:       d.Get("region").(string),
		InstanceType: d.Get("instance_type").(string),
		ChiaVersion:  d.Get("chia_version").(string),
		Network:      d.Get("network").(string),
	})

	if err != nil {
		return diag.Errorf("There was an error creating the node")
	}

	d.SetId(result.NodeId)

	tflog.Trace(ctx, "created marmotcore node")

	return diags
}

func nodeMarmotCoreRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func nodeMarmotCoreUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func nodeMarmotCoreDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}
