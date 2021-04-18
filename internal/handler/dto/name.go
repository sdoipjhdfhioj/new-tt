package dto

type SetNameRequest struct {
	ID   string
	Name string
}

type NameResponse struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
