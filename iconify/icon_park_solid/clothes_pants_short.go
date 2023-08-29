package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesPantsShort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSClothesPantsShort0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="m6 36l2-24h32l2 24H26.842L24 25l-2.842 11H6Z"/><path stroke="#000" d="m24 12l3 7m-3-7l-4 7.5"/><path stroke="#fff" d="M18 12h12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSClothesPantsShort0)"/>`),
		g.Group(children),
	)
}