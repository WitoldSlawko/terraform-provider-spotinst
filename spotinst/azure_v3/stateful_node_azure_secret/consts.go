package stateful_node_azure_secret

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/tree/main/spotinst/commons"

const (
	Secret commons.FieldName = "secret"

	SourceVault       commons.FieldName = "source_vault"
	Name              commons.FieldName = "name"
	ResourceGroupName commons.FieldName = "resource_group_name"

	VaultCertificates commons.FieldName = "vault_certificates"
	CertificateURL    commons.FieldName = "certificate_url"
	CertificateStore  commons.FieldName = "certificate_store"
)
