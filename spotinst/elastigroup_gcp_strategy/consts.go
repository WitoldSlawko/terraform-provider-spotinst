package elastigroup_gcp_strategy

import "github.com/WitoldSlawko/terraform-provider-spotinst/spotinst/commons"

const (
	PreemptiblePercentage commons.FieldName = "preemptible_percentage"
	OnDemandCount         commons.FieldName = "ondemand_count"
	DrainingTimeout       commons.FieldName = "draining_timeout"
	FallbackToOnDemand    commons.FieldName = "fallback_to_ondemand"
	ProvisioningModel     commons.FieldName = "provisioning_model"
)
