package ocean_aws_strategy

import (
	"https://github.com/WitoldSlawko/terraform-provider-spotinst/tree/main/spotinst/commons"
)

const (
	SpotPercentage           commons.FieldName = "spot_percentage"
	FallbackToOnDemand       commons.FieldName = "fallback_to_ondemand"
	UtilizeReservedInstances commons.FieldName = "utilize_reserved_instances"
	DrainingTimeout          commons.FieldName = "draining_timeout"
	GracePeriod              commons.FieldName = "grace_period"
	UtilizeCommitments       commons.FieldName = "utilize_commitments"
)
