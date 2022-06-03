package mrscaler_aws_strategy

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/tree/main/spotinst/commons"

const (
	Clone = "clone"
	Wrap  = "wrap"
	New   = "new"
)

const (
	Strategy commons.FieldName = "strategy"

	ProvisioningTimeout commons.FieldName = "provisioning_timeout"
	TimeoutAction       commons.FieldName = "timeout_action"
	Timeout             commons.FieldName = "timeout"
	ReleaseLabel        commons.FieldName = "release_label"
	Retries             commons.FieldName = "retries"
)
