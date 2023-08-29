package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-opacity=".3" d="M8 8h8v7h3V5H5v10h3V8Z"/><path d="M5 17h14v3H5v-3Z"/></g>`),
		g.Group(children),
	)
}