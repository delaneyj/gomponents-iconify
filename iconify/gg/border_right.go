package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-opacity=".3" d="M8 16V8h7V5H5v14h10v-3H8Z"/><path d="M17 19V5h3v14h-3Z"/></g>`),
		g.Group(children),
	)
}