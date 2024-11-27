package provider

import (
	"context"
	"os/exec"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &shelleDataSource{}
	_ datasource.DataSourceWithConfigure = &shelleDataSource{}
)

func NewShelleDataSource() datasource.DataSource {
	return &shelleDataSource{}
}

// shelleDataSource is the data source implementation.
type shelleDataSource struct {
}

// shelleDataSourceModel maps the data source schema data.
type shelleDataSourceModel struct {
	CommandOutput types.String `tfsdk:"command_output"`
	CommandText   types.String `tfsdk:"command_text"`
}

// Metadata returns the data source type name.
func (d *shelleDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_shelle"
}

// Schema defines the schema for the data source.
func (d *shelleDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"command_text": schema.StringAttribute{
				Required: true,
			},
			"command_output": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *shelleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state shelleDataSourceModel

	// Get the command_text attribute from the request
	var passedIn shelleDataSourceModel
	req.Config.Get(ctx, &passedIn)
	state.CommandText = passedIn.CommandText

	// Execute the command_text using /bin/sh
	cmd := exec.CommandContext(ctx, "/bin/sh", "-c", state.CommandText.ValueString())

	// Capture the output
	output, err := cmd.Output()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error executing command",
			"Could not execute command: "+err.Error(),
		)
		return
	}

	// Set the command_output attribute in the state
	state.CommandOutput = types.StringValue(string(output))

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Configure adds the provider configured client to the data source.
func (d *shelleDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {

}
