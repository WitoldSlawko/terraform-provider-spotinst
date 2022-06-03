package elastigroup_azure_launch_configuration

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/tree/main/spotinst/commons"

const (
	Prefix = "azure_launch_configuration_"
)

const (
	UserData                 commons.FieldName = "user_data"
	ShutdownScript           commons.FieldName = "shutdown_script"
	CustomData               commons.FieldName = "custom_data"
	ManagedServiceIdentities commons.FieldName = "managed_service_identities"

	ResourceGroupName commons.FieldName = "resource_group_name"
	Name              commons.FieldName = "name"
)
