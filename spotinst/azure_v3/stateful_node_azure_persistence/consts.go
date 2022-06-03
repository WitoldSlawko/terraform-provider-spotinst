package stateful_node_azure_persistence

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/tree/main/spotinst/commons"

const (
	ShouldPersistOSDisk      commons.FieldName = "should_persist_os_disk"
	OSDiskPersistenceMode    commons.FieldName = "os_disk_persistence_mode"
	ShouldPersistDataDisks   commons.FieldName = "should_persist_data_disks"
	DataDisksPersistenceMode commons.FieldName = "data_disks_persistence_mode"
	ShouldPersistNetwork     commons.FieldName = "should_persist_network"
	ShouldPersistVM          commons.FieldName = "should_persist_vm"
)
