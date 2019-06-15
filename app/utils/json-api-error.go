package utils

// JSONAPIError ==> replaces the type from the old gson api package
// This can be improved, but prevents us from having to refactor many lines
type JSONAPIError struct {
	ID     string `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
	Code   string `json:"code,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
	Source string `json:"source,omitempty"`
}
