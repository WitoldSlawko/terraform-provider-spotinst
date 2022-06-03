package managed_instance_aws

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/commons"

const (
	Name        commons.FieldName = "name"
	Description commons.FieldName = "description"
	Region      commons.FieldName = "region"

	// - Instance Action ----------------------
	ManagedInstanceAction commons.FieldName = "managed_instance_action"
	ActionType            commons.FieldName = "type"
	// ----------------------------------------
)
