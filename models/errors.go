package models

type Error struct {
	Code    int    `json:"code,omitempty"`
	Status  string `json:"status,omitempty"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}

type Errors struct {
	Code    int     `json:"code"`
	Status  string  `json:"status"`
	Message string  `json:"message,omitempty"`
	Errors  []Error `json:"errors"`
}
