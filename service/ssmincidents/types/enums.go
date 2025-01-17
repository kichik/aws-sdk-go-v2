// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type IncidentRecordStatus string

// Enum values for IncidentRecordStatus
const (
	IncidentRecordStatusOpen     IncidentRecordStatus = "OPEN"
	IncidentRecordStatusResolved IncidentRecordStatus = "RESOLVED"
)

// Values returns all known values for IncidentRecordStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IncidentRecordStatus) Values() []IncidentRecordStatus {
	return []IncidentRecordStatus{
		"OPEN",
		"RESOLVED",
	}
}

type ItemType string

// Enum values for ItemType
const (
	ItemTypeAnalysis   ItemType = "ANALYSIS"
	ItemTypeIncident   ItemType = "INCIDENT"
	ItemTypeMetric     ItemType = "METRIC"
	ItemTypeParent     ItemType = "PARENT"
	ItemTypeAttachment ItemType = "ATTACHMENT"
	ItemTypeOther      ItemType = "OTHER"
	ItemTypeAutomation ItemType = "AUTOMATION"
)

// Values returns all known values for ItemType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (ItemType) Values() []ItemType {
	return []ItemType{
		"ANALYSIS",
		"INCIDENT",
		"METRIC",
		"PARENT",
		"ATTACHMENT",
		"OTHER",
		"AUTOMATION",
	}
}

type RegionStatus string

// Enum values for RegionStatus
const (
	// All operations have completed successfully and the region is ready to use
	RegionStatusActive RegionStatus = "ACTIVE"
	// The region is in the process of being created.
	RegionStatusCreating RegionStatus = "CREATING"
	// The region is in the process of being deleted.
	RegionStatusDeleting RegionStatus = "DELETING"
	// The region is not healthy and we cannot automatically fix it.
	RegionStatusFailed RegionStatus = "FAILED"
)

// Values returns all known values for RegionStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (RegionStatus) Values() []RegionStatus {
	return []RegionStatus{
		"ACTIVE",
		"CREATING",
		"DELETING",
		"FAILED",
	}
}

type ReplicationSetStatus string

// Enum values for ReplicationSetStatus
const (
	// All operations have completed successfully and the replication set is ready to
	// use
	ReplicationSetStatusActive ReplicationSetStatus = "ACTIVE"
	// Replication set is in the process of being created.
	ReplicationSetStatusCreating ReplicationSetStatus = "CREATING"
	// Replication set is in the process of being updated.
	ReplicationSetStatusUpdating ReplicationSetStatus = "UPDATING"
	// Replication set is in the process of being deleted.
	ReplicationSetStatusDeleting ReplicationSetStatus = "DELETING"
	// Replication set is not healthy and we cannot fix it.
	ReplicationSetStatusFailed ReplicationSetStatus = "FAILED"
)

// Values returns all known values for ReplicationSetStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReplicationSetStatus) Values() []ReplicationSetStatus {
	return []ReplicationSetStatus{
		"ACTIVE",
		"CREATING",
		"UPDATING",
		"DELETING",
		"FAILED",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeResponsePlan   ResourceType = "RESPONSE_PLAN"
	ResourceTypeIncidentRecord ResourceType = "INCIDENT_RECORD"
	ResourceTypeTimelineEvent  ResourceType = "TIMELINE_EVENT"
	ResourceTypeReplicationSet ResourceType = "REPLICATION_SET"
	ResourceTypeResourcePolicy ResourceType = "RESOURCE_POLICY"
)

// Values returns all known values for ResourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"RESPONSE_PLAN",
		"INCIDENT_RECORD",
		"TIMELINE_EVENT",
		"REPLICATION_SET",
		"RESOURCE_POLICY",
	}
}

type ServiceCode string

// Enum values for ServiceCode
const (
	ServiceCodeSsmIncidents ServiceCode = "ssm-incidents"
)

// Values returns all known values for ServiceCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ServiceCode) Values() []ServiceCode {
	return []ServiceCode{
		"ssm-incidents",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAscending  SortOrder = "ASCENDING"
	SortOrderDescending SortOrder = "DESCENDING"
)

// Values returns all known values for SortOrder. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"ASCENDING",
		"DESCENDING",
	}
}

type SsmTargetAccount string

// Enum values for SsmTargetAccount
const (
	SsmTargetAccountResponsePlanOwnerAccount SsmTargetAccount = "RESPONSE_PLAN_OWNER_ACCOUNT"
	SsmTargetAccountImpactedAccount          SsmTargetAccount = "IMPACTED_ACCOUNT"
)

// Values returns all known values for SsmTargetAccount. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SsmTargetAccount) Values() []SsmTargetAccount {
	return []SsmTargetAccount{
		"RESPONSE_PLAN_OWNER_ACCOUNT",
		"IMPACTED_ACCOUNT",
	}
}

type TimelineEventSort string

// Enum values for TimelineEventSort
const (
	TimelineEventSortEventTime TimelineEventSort = "EVENT_TIME"
)

// Values returns all known values for TimelineEventSort. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TimelineEventSort) Values() []TimelineEventSort {
	return []TimelineEventSort{
		"EVENT_TIME",
	}
}
