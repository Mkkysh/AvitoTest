package dto

type ChangeSegments struct {
	AddSegments    []interface{} `json:"add_segments"`
	RemoveSegments []interface{} `json:"remove_segments"`
}
