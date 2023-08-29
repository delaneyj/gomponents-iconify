package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScoreSheets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.5 30.867V6.5a2 2 0 0 0-2-2h-27a2 2 0 0 0-2 2v35a2 2 0 0 0 2 2h27a2 2 0 0 0 2-2v-1.469"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.134 37.66L21.288 21.775V17.25h4.632l15.785 15.828m2.61 4.907a1.623 1.623 0 0 0 0-2.291l-2.61-2.616l-4.57 4.582l2.61 2.616a1.614 1.614 0 0 0 2.284 0ZM13 10.5h22m-22 6.75h8.288M32.652 24H35m-22 0h10.508M13 30.75h16.989M13 37.5h22"/>`),
		g.Group(children),
	)
}