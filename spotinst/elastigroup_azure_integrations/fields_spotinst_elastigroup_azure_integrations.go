package elastigroup_azure_integrations

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/commons"

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	SetupKubernetes(fieldsMap)
	SetupMultaiRuntime(fieldsMap)
}
