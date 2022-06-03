package elastigroup_azure_login

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/tree/main/spotinst/commons"

const (
	Prefix = "azure_login_"
)

const (
	Login        commons.FieldName = "login"
	UserName     commons.FieldName = "user_name"
	SSHPublicKey commons.FieldName = "ssh_public_key"
	Password     commons.FieldName = "password"
)
