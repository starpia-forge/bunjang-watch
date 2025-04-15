package filter

type FilterConfig struct {
	Keywords     []string `json:"keywords,omitempty"`
	MinimumPrice int      `json:"minimum_price,omitempty"`
	MaximumPrice int      `json:"maximum_price,omitempty"`
	IncludeUsed  bool     `json:"include_used,omitempty"`
}
