package multai_routing_rule

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/commons"

const (
	BalancerID    commons.FieldName = "balancer_id"
	ListenerID    commons.FieldName = "listener_id"
	Route         commons.FieldName = "route"
	Priority      commons.FieldName = "priority"
	Strategy      commons.FieldName = "strategy"
	MiddlewareIDs commons.FieldName = "middleware_ids"
	TargetSetIDs  commons.FieldName = "target_set_ids"
	Tags          commons.FieldName = "tags"

	TagKey   commons.FieldName = "key"
	TagValue commons.FieldName = "value"
)
