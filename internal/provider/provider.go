package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &shelleProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &shelleProvider{
			version: version,
		}
	}
}

// shelleProvider is the provider implementation.
type shelleProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// Metadata returns the provider type name.
func (p *shelleProvider) Metadata(ctx context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	tflog.Warn(ctx, "Hello from provider Metadata")
	resp.TypeName = "shelle"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *shelleProvider) Schema(ctx context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	tflog.Warn(ctx, "Hello from provider Schema")
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{},
	}
}

func (p *shelleProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Warn(ctx, "Hello from provider configure")
}

// DataSources defines the data sources implemented in the provider.
func (p *shelleProvider) DataSources(c context.Context) []func() datasource.DataSource {
	tflog.Warn(c, "Hello from provider DataSources")
	return []func() datasource.DataSource{
		NewShelleDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *shelleProvider) Resources(c context.Context) []func() resource.Resource {
	tflog.Error(c, "Hello from provider Resources")
	return []func() resource.Resource{}
}
