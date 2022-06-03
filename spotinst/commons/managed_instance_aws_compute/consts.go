package managed_instance_aws_compute

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/tree/main/spotinst/commons"

const (
	SubnetIDs commons.FieldName = "subnet_ids"
	VpcID     commons.FieldName = "vpc_id"
	ElasticIP commons.FieldName = "elastic_ip"
	PrivateIP commons.FieldName = "private_ip"
)
