package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PathTrim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M5 5h10v3H8v7H5V5Z"/><path d="M19 9H9v10h10V9Z"/></g>`),
		g.Group(children),
	)
}