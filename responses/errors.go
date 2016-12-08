package responses

type Error struct {
	Status  int    `json:"status,omitempty"`
	Title   string `json:"title,omitempty"`
	Message string `json:"message"`
}

type Errors struct {
	Status  int     `json:"status"`
	Title   string  `json:"title"`
	Message string  `json:"message,omitempty"`
	Errors  []Error `json:"errors"`
}
