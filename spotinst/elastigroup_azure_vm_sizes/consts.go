package elastigroup_azure_vm_sizes

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/tree/main/spotinst/commons"

const (
	Prefix = "azure_vm_sizes_"
)

const (
	VMSizes     commons.FieldName = "elastigroup_azure_vm_sizes"
	OnDemand    commons.FieldName = "od_sizes"
	LowPriority commons.FieldName = "low_priority_sizes"
)
