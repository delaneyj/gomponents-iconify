package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m2 8l1.304 1.043a4 4 0 0 0 5.995-1.181L12 3l2.701 4.862a4 4 0 0 0 5.995 1.18L22 8l-1.754 8.77a2.564 2.564 0 0 1-1.367 1.79v0a15.381 15.381 0 0 1-13.758 0v0a2.564 2.564 0 0 1-1.367-1.79L2 8Z"/><path d="M8 15c2.596 1.333 5.404 1.333 8 0"/></g>`),
		g.Group(children),
	)
}