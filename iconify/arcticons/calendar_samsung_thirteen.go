package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSamsungThirteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.932 18.925l3.5-1.925m0 0v14m9.179-7a3.467 3.467 0 0 1 3.457 3.457h0a3.467 3.467 0 0 1-3.457 3.457h-1.383c-2.42 0-3.284-.346-4.32-1.21m-.001-11.408c1.037-.864 2.074-1.21 4.321-1.21h1.383a3.467 3.467 0 0 1 3.457 3.457h0A3.467 3.467 0 0 1 27.61 24h-3.457"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.5 5.5h-4v1a3.5 3.5 0 1 1-7 0v-1h-8v1a3.5 3.5 0 1 1-7 0v-1h-3a4 4 0 0 0-4 4v29a4 4 0 0 0 4 4h29a4 4 0 0 0 4-4v-29a4 4 0 0 0-4-4Z"/>`),
		g.Group(children),
	)
}