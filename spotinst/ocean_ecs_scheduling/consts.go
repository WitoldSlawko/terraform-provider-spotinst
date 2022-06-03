package ocean_ecs_scheduling

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/tree/main/spotinst/commons"

const (
	ScheduledTask          commons.FieldName = "scheduled_task"
	ShutdownHours          commons.FieldName = "shutdown_hours"
	TimeWindows            commons.FieldName = "time_windows"
	ShutdownHoursIsEnabled commons.FieldName = "is_enabled"
	Tasks                  commons.FieldName = "tasks"
	TasksIsEnabled         commons.FieldName = "is_enabled"
	CronExpression         commons.FieldName = "cron_expression"
	TaskType               commons.FieldName = "task_type"
)
