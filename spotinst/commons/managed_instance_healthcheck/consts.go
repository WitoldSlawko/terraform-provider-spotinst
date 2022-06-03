package managed_instance_healthcheck

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/tree/main/spotinst/commons"

const (
	HealthCheckType   commons.FieldName = "health_check_type"
	AutoHealing       commons.FieldName = "auto_healing"
	GracePeriod       commons.FieldName = "grace_period"
	UnhealthyDuration commons.FieldName = "unhealthy_duration"
)
