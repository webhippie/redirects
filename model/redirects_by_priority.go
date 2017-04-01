package model

// RedirectsByPriority represents a collection of redirects for sorting.
type RedirectsByPriority []*Redirect

// Len just implements the default sorting interface.
func (s RedirectsByPriority) Len() int {
	return len(s)
}

// Swap just implements the default sorting interface.
func (s RedirectsByPriority) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less just implements the default sorting interface.
func (s RedirectsByPriority) Less(i, j int) bool {
	return s[i].Priority < s[j].Priority
}
