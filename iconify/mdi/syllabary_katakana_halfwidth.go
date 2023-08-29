package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SyllabaryKatakanaHalfwidth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 4v2h6l-1 4.81l2 .39l1-5.1V4m-6 5v3c0 2.86-.66 5.29-1.92 6.61L9.5 20c1.85-1.92 2.5-4.85 2.5-8V9Z"/>`),
		g.Group(children),
	)
}