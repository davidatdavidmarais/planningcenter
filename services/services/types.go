package services

type SchedulesResponse struct {
	Data []Schedule `json:"data"`
}

type Schedule struct {
	ID            string              `json:"id"`
	Type          string              `json:"type"`
	Relationships RelationshipDetails `json:"relationships"`
}

type RelationshipDetails struct {
	ServiceType RelationshipType `json:"service_type"`
	Plan        RelationshipType `json:"plan"`
}

type RelationshipType struct {
	Data RelationshipData `json:"data"`
}

type RelationshipData struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}
