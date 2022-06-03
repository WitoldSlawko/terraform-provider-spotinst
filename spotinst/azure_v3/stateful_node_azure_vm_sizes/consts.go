package stateful_node_azure_vm_sizes

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/tree/main/spotinst/commons"

const (
	OnDemandSizes      commons.FieldName = "od_sizes"
	SpotSizes          commons.FieldName = "spot_sizes"
	PreferredSpotSizes commons.FieldName = "preferred_spot_sizes"
)
