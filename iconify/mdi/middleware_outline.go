package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MiddlewareOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m23 12l-4-4v3h-3.14c-.45-1.72-2-3-3.86-3s-3.41 1.28-3.86 3H5L2 8v8l3-3h3.14c.45 1.72 2 3 3.86 3s3.41-1.28 3.86-3H19v3l4-4m-11 2c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2Z"/>`),
		g.Group(children),
	)
}