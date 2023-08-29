package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirBike(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAirBike0"><g fill="none" stroke-width="4"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M4 44h40M28 30h6.19c2.27 0 6.81 1.344 6.81 6.72V44m-6-14l5-11l-6-13m-5 2l10-4"/><circle cx="20" cy="30" r="8" fill="#fff" stroke="#fff"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M20 30h8"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="m21 22l-7-9m-4 0h8"/><path stroke="#fff" d="M20 38a8 8 0 1 0 0-16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAirBike0)"/>`),
		g.Group(children),
	)
}