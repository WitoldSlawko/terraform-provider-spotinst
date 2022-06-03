package stateful_node_azure_load_balancer

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/tree/main/spotinst/commons"

const (
	LoadBalancer commons.FieldName = "load_balancer"

	Type              commons.FieldName = "type"
	Name              commons.FieldName = "name"
	ResourceGroupName commons.FieldName = "resource_group_name"
	SKU               commons.FieldName = "sku"
	BackendPoolNames  commons.FieldName = "backend_pool_names"
)
