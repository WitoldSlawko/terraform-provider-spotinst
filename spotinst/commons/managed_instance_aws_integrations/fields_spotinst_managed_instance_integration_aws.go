package managed_instance_aws_integrations

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/commons"

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	SetupRoute53(fieldsMap)
	SetupLoadBalancers(fieldsMap)
}
