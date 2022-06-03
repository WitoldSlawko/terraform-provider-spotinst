package ocean_aks_virtual_node_group_launch_specification

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/tree/main/spotinst/commons"

const (
	LaunchSpecification commons.FieldName = "launch_specification"
)

const (
	OSDisk                  commons.FieldName = "os_disk"
	SizeGB                  commons.FieldName = "size_gb"
	Type                    commons.FieldName = "type"
	UtilizeEphemeralStorage commons.FieldName = "utilize_ephemeral_storage"
)

const (
	Tag      commons.FieldName = "tag"
	TagKey   commons.FieldName = "key"
	TagValue commons.FieldName = "value"
)
