package ocean_gke_launch_spec_import

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/tree/main/spotinst/commons"

type LabelField string
type MetadataField string

const (
	OceanId      commons.FieldName = "ocean_id"
	NodePoolName commons.FieldName = "node_pool_name"
)
