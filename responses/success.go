package responses

type Pages struct {
	Current int `json:"current,omitempty"`
	Next    int `json:"next,omitempty"`
	Last    int `json:"last,omitempty"`
	Limit   int `json:"limit,omitempty"`
	Total   int `json:"total,omitempty"`
}

type Links struct {
	Self     string `json:"self"`
	First    string `json:"first,omitempty"`
	Previous string `json:"prev,omitempty"`
	Next     string `json:"next,omitempty"`
	Last     string `json:"last,omitempty"`
}

type Success struct {
	Status int                    `json:"status"`
	Title  string                 `json:"title"`
	Meta   map[string]interface{} `json:"meta,omitempty"`
	Links  Links                  `json:"links"`
	Data   interface{}            `json:"data"`
}
