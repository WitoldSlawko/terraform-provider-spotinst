package elastigroup_azure_integrations

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/tree/main/spotinst/commons"

const (
	// - KUBERNETES ----------------------
	IntegrationKubernetes commons.FieldName = "integration_kubernetes"
	ClusterIdentifier     commons.FieldName = "cluster_identifier"
	// -----------------------------------

	// - MULTAI-RUNTIME ------------------
	IntegrationMultaiRuntime commons.FieldName = "integration_multai_runtime"
	DeploymentId             commons.FieldName = "deployment_id"
	// -----------------------------------
)
