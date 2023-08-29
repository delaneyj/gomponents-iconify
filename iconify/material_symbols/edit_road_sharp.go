package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditRoadSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 11.9V4h2v5.9l-2 2ZM4 20V4h2v16H4Zm6-12V4h2v4h-2Zm0 6v-4h2v4h-2Zm0 6v-4h2v4h-2Zm4 0v-2.125l5.075-5.1l2.15 2.1l-5.1 5.125H14Zm7.95-5.825l-2.125-2.125l1.425-1.425l2.1 2.15l-1.4 1.4Z"/>`),
		g.Group(children),
	)
}