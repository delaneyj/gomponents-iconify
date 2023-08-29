package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LevelFiveTitle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 8v32M24 8v32M7 24h16m17-2.99h-8v7.024C32 28 34 27 37 27s4 2.534 4 6.5s-1 6.5-5 6.5c-3 0-4-2-4-3.992"/>`),
		g.Group(children),
	)
}