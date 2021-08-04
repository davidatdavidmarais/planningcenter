package services

// Common

type RelationshipType struct {
	Data RelationshipData `json:"data"`
}

type RelationshipData struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type Base struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// Schedules

type SchedulesResponse struct {
	Data []Schedule `json:"data"`
}

type Schedule struct {
	Base

	Relationships ScheduleRelationships `json:"relationships"`
}

type ScheduleRelationships struct {
	ServiceType RelationshipType `json:"service_type"`
	Plan        RelationshipType `json:"plan"`
}

// Items

type ItemsResponse struct {
	Data []Item `json:"data"`
}

type Item struct {
	Base

	Attributes ItemAttributes `json:"attributes"`

	Relationships ItemRelationships `json:"relationships"`
}

type ItemAttributes struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Sequence    int32  `json:"sequence"`
	ItemType    string `json:"item_type"`
}

type ItemRelationships struct {
	Arrangement RelationshipType `json:"arrangement"`
}
