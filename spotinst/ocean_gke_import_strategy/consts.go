package ocean_gke_import_strategy

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/tree/main/spotinst/commons"

const (
	Strategy              commons.FieldName = "strategy"
	DrainingTimeout       commons.FieldName = "draining_timeout"
	ProvisioningModel     commons.FieldName = "provisioning_model"
	PreemptiblePercentage commons.FieldName = "preemptible_percentage"
)
