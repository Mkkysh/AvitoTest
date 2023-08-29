package dto

type LogResponse struct {
	IdUser      int    `json:"id_user"`
	NameSegment string `json:"name_segment"`
	Operation   string `json:"operation"`
	Timestamp   string `json:"timestamp"`
}
