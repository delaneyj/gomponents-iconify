package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PartyBalloon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPartyBalloon0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M34 16c1-8-2.873-12-9.873-12C17.127 4 13 9 14 16s7.255 12 10.127 12C27 28 33 24 34 16Z"/><path d="M25 28c-2 .97-5 3.889-5 7s10 1.444 10 4.556C30 42.666 19 44 19 44"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPartyBalloon0)"/>`),
		g.Group(children),
	)
}