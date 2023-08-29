package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-opacity=".3" d="M16 8v8H9v3h10V5H9v3h7Z"/><path d="M7 5v14H4V5h3Z"/></g>`),
		g.Group(children),
	)
}