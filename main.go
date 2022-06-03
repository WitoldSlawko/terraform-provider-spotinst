package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"https://github.com/WitoldSlawko/terraform-provider-spotinst"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: spotinst.Provider})
}
