package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SyllabaryHangul(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 4v2H4v2h2.39C5.55 8.74 5 9.8 5 11c0 2.2 1.8 4 4 4s4-1.8 4-4c0-1.2-.55-2.26-1.39-3H14V6h-4V4m5 0v12h2v-5h3V9h-3V4M9 9c1.12 0 2 .88 2 2s-.88 2-2 2s-2-.88-2-2s.88-2 2-2m-2 7v4h10v-2H9v-2Z"/>`),
		g.Group(children),
	)
}