package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Seat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 18v3h3v-3h10v3h3v-6H4v3m15-8h3v3h-3v-3M2 10h3v3H2v-3m15 3H7V5a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v8Z"/>`),
		g.Group(children),
	)
}