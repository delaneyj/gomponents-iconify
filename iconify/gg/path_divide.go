package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PathDivide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M5 5h10v4H9v6H5V5Z"/><path d="M9 15v4h10V9h-4v6H9Z"/><path d="M10 10h4v4h-4v-4Z"/></g>`),
		g.Group(children),
	)
}