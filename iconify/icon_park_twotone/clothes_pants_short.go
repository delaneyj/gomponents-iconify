package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesPantsShort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTClothesPantsShort0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="m6 36l2-24h32l2 24H26.842L24 25l-2.842 11H6Z"/><path d="m24 12l3 7m-3-7l-4 7.5M18 12h12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTClothesPantsShort0)"/>`),
		g.Group(children),
	)
}