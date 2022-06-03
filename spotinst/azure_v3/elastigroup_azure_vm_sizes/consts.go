package elastigroup_azure_vm_sizes

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/tree/main/spotinst/commons"

const (
	Prefix = "azure_vm_sizes_"
)

const (
	OnDemandSizes commons.FieldName = "od_sizes"
	SpotSizes     commons.FieldName = "spot_sizes"
)
