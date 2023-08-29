package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabySling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBabySling0"><g fill="none" stroke="#fff" stroke-width="4"><circle cx="24" cy="10" r="5" fill="#555"/><path fill="#555" stroke-linecap="round" stroke-linejoin="round" d="M24 21C14 21 11 6 11 6L6 8l3 16.5c1.167.667 5.5 3 8 6.5s2 7 7.5 7s6-4.5 7.5-7s5.667-6 7-6.5L42 8l-5-2s-3 15-13 15Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M34 29s7 4 7 13h-5c0-6-5-9-5-9m-16-4s-7 4-7 13h5c0-6 5-9 5-9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBabySling0)"/>`),
		g.Group(children),
	)
}