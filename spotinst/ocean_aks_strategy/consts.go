package ocean_aks_strategy

import (
	"github.com/WitoldSlawko/terraform-provider-spotinst/spotinst/commons"
)

const (
	Strategy           commons.FieldName = "strategy"
	SpotPercentage     commons.FieldName = "spot_percentage"
	FallbackToOnDemand commons.FieldName = "fallback_to_ondemand"
)
