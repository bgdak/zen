package zen

import "strings"

// parseParams parse url parameters
func (r *route) parseParams(c *Context) {
	pattern := c.Req.URL.Path
	pattern = strings.TrimSuffix(strings.TrimPrefix(pattern, "/"), "/")
	parts := strings.Split(pattern, "/")

	for i, k := range r.params {
		c.params[k] = parts[i]
	}
}