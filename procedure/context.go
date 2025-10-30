package procedure

type Context struct {
	// When executing a procedure, these are the path variables extracted from the path if any
	// For example, if the procedure path is "/user/{id}/profile" and the actual path is "/user/123/profile",
	// then pathParams will contain map[string]string{"id": "123"}
	pathParams map[string]string
}

func (c *Context) Param(name string) string {
	if c.pathParams == nil {
		return ""
	}

	if val, ok := c.pathParams[name]; ok {
		return val
	}

	return ""
}
