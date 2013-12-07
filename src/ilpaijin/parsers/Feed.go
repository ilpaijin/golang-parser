package parsers

type Feed interface {
	Parse(parseType string)
	Csvize()
}
