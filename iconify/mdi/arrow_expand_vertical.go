package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowExpandVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 9v6h3l-4 4l-4-4h3V9H8l4-4l4 4h-3M4 2h16v2H4V2m0 18h16v2H4v-2Z"/>`),
		g.Group(children),
	)
}