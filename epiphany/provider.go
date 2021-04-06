package epiphany

import (
	"context"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"example_cluster": resourceEpiphanyCluster(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	id, _ := uuid.NewUUID()
	runner := Runner{
		id:   id,
		name: "test runner instance",
	}

	var diagnostics diag.Diagnostics

	diagnostics = append(diagnostics, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Unable to create HashiCups client",
		Detail:   "Unable to auth user for authenticated HashiCups client",
	})

	return runner, diagnostics
}

type Runner struct {
	id   uuid.UUID
	name string
}
