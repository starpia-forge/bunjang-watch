package filter

type FilterConfig struct {
	Keywords     []string
	MinimumPrice int
	MaximumPrice int
	IncludeUsed  bool
}
