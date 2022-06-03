package subscription

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/tree/main/spotinst/commons"

const (
	Prefix = "subscription_"
)

const (
	ResourceId commons.FieldName = "resource_id"
	EventType  commons.FieldName = "event_type"
	Protocol   commons.FieldName = "protocol"
	Endpoint   commons.FieldName = "endpoint"
	Format     commons.FieldName = "format"
)
