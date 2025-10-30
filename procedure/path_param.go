package procedure

// TODO: move this to a separate package: api? router?
type pathParam struct {
	name    string
	value   string
	pattern string
}

func NewPathParam(name, value, pattern string) *pathParam {
	return &pathParam{name: name, value: value, pattern: pattern}
}

func (pp *pathParam) Name() string  { return pp.name }
func (pp *pathParam) Value() string { return pp.value }

func (pp *pathParam) Match(value string) bool {
	panic("unimplemented")
}
