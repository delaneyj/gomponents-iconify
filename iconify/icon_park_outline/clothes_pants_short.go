package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesPantsShort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m6 36l2-24h32l2 24H26.842L24 25l-2.842 11H6Zm18-24l3 7m-3-7l-4 7.5M18 12h12"/>`),
		g.Group(children),
	)
}