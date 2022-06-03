package ocean_aks_launch_specification

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/commons"

const (
	CustomData        commons.FieldName = "custom_data"
	ResourceGroupName commons.FieldName = "resource_group_name"
)

const (
	ManagedServiceIdentity                  commons.FieldName = "managed_service_identity"
	ManagedServiceIdentityResourceGroupName commons.FieldName = "resource_group_name"
	ManagedServiceIdentityName              commons.FieldName = "name"
)

const (
	Tag      commons.FieldName = "tag"
	TagKey   commons.FieldName = "key"
	TagValue commons.FieldName = "value"
)
