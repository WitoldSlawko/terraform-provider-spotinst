package elastigroup_azure_health_check

import "github.com/WitoldSlawko/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "health_check"
)

const (
	HealthCheck     commons.FieldName = "health_check"
	AutoHealing     commons.FieldName = "auto_healing"
	HealthCheckType commons.FieldName = "health_check_type"
	GracePeriod     commons.FieldName = "grace_period"
)
