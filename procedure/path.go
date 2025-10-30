package procedure

type Path struct {
	rawPath string
	params  map[string]*pathParam
}

// TODO: fuck, we're building a router again
func NewPath(path string) (*Path, error) {
	p := &Path{rawPath: normalizePath(path)}

	// TODO: implement
	panic("unimplemented")

	return p, nil
}

func (p *Path) Parse() error {
	panic("unimplemented")
}

// Matches checks if the provided route matches this route, it checks that all params match the required pattern, if any
func (p *Path) Matches(path string) {
	panic("unimplemented")
}

func (p *Path) Path() string {
	return p.rawPath
}

func normalizePath(path string) string {
	if path == "" {
		return "/"
	}

	if path[0] != '/' {
		path = "/" + path
	}

	if len(path) > 1 && path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	return path
}
