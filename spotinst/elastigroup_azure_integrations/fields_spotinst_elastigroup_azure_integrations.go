package elastigroup_azure_integrations

import "github.com/WitoldSlawko/terraform-provider-spotinst/spotinst/commons"

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	SetupKubernetes(fieldsMap)
	SetupMultaiRuntime(fieldsMap)
}
