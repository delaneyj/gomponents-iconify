package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlightAirflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFlightAirflow0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path d="M6 25c0-9.941 8.283-18 18.5-18S43 15.059 43 25"/><path fill="#555" stroke-linejoin="round" d="m10 35l-1.064-5s-3.435 3.109-4.58 5.74C3.211 38.37 4.852 41 8 41h36l-8-5.978L10 35Z"/><path stroke-linejoin="round" d="M29 35L18 25h-3l2 10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFlightAirflow0)"/>`),
		g.Group(children),
	)
}