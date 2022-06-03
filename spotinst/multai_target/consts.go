package multai_target

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/tree/main/spotinst/commons"

const (
	BalancerID  commons.FieldName = "balancer_id"
	TargetSetID commons.FieldName = "target_set_id"
	Name        commons.FieldName = "name"
	Host        commons.FieldName = "host"
	Port        commons.FieldName = "port"
	Weight      commons.FieldName = "weight"
	Tags        commons.FieldName = "tags"

	TagKey   commons.FieldName = "key"
	TagValue commons.FieldName = "value"
)
