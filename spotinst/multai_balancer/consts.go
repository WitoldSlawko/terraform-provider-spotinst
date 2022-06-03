package multai_balancer

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/tree/main/spotinst/commons"

const (
	Name               commons.FieldName = "name"
	Scheme             commons.FieldName = "scheme"
	DNSCnameAliases    commons.FieldName = "dns_cname_aliases"
	ConnectionTimeouts commons.FieldName = "connection_timeouts"
	Tags               commons.FieldName = "tags"

	Idle     commons.FieldName = "idle"
	Draining commons.FieldName = "draining"

	TagKey   commons.FieldName = "key"
	TagValue commons.FieldName = "value"
)
