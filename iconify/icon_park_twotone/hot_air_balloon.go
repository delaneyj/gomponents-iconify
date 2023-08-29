package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HotAirBalloon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHotAirBalloon0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M39 18.352C39 27.855 31 36 24 36S9 27.855 9 18.352C9 10.654 14.893 4 24 4s15 6.654 15 14.352Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="M25 4c3.7 3.819 7 10.12 7 16c0 5.807-3.38 12.192-7 16M23 4c-4.317 4.087-7 9.706-7 16c0 6.215 2.777 11.924 7 16"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M22 16.018s5.065 1.017 9 0C34.935 15 38 13 38 13M9 20s5 2.5 9 3.5s8 .5 8 .5"/><path fill="#555" d="m18 35l1.4 7.095S22.125 44 24 44c1.875 0 4.6-1.905 4.6-1.905l1.394-7.065L24 36l-6-1Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m30 35l-.006.03m0 0L28.6 42.095S25.875 44 24 44c-1.875 0-4.6-1.905-4.6-1.905L18 35l6 1l5.994-.97Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHotAirBalloon0)"/>`),
		g.Group(children),
	)
}