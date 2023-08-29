package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SeatOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 5v7H9V5h6m0-2H9a2 2 0 0 0-2 2v9h10V5a2 2 0 0 0-2-2m7 7h-3v3h3v-3M5 10H2v3h3v-3m15 5H4v6h2v-4h12v4h2v-6Z"/>`),
		g.Group(children),
	)
}