package services

import "time"

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

	Attributes ScheduleAttributes `json:"attributes"`

	Relationships ScheduleRelationships `json:"relationships"`
}

type ScheduleAttributes struct {
	OrganizationName string    `json:"organization_name"`
	ServiceTypeName  string    `json:"service_type_name"`
	DateTime         time.Time `json:"sort_date"`
	TeamName         string    `json:"team_name"`
	TeamPositionName string    `json:"team_position_name"`
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
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Sequence    int32   `json:"sequence"`
	ItemType    string  `json:"item_type"`
	Length      float32 `json:"length"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ItemRelationships struct {
	Arrangement RelationshipType `json:"arrangement"`
}

// Arrangement

type ArrangementResponse struct {
	Data Arrangement `json:"data"`
}

type Arrangement struct {
	Base

	Attributes ArrangementAttributes `json:"attributes"`
}

type ArrangementAttributes struct {
	BPM           float32 `json:"bpm"`
	ChordChart    string  `json:"chord_chart"`
	ChordChartKey string  `json:"chord_chart_key"`
	Lyrics        string  `json:"lyrics"`
	Meter         string  `json:"meter"`
	Name          string  `json:"name"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
