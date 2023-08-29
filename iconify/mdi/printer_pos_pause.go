package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterPosPause(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 9H7V4h10v5m3 4.09c-.33-.05-.66-.09-1-.09c-3.31 0-6 2.69-6 6H4v-7a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v1.09M10 12H6v2h4v-2m6 10h2v-6h-2v6m4-6v6h2v-6h-2Z"/>`),
		g.Group(children),
	)
}