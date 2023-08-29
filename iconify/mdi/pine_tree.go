package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PineTree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 21v-3H3l5-5H5l5-5H7l5-5l5 5h-3l5 5h-3l5 5h-7v3h-4Z"/>`),
		g.Group(children),
	)
}