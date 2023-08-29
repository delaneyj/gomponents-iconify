package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplaneWindowOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAirplaneWindowOne0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path stroke="#fff" stroke-linecap="round" d="m30.349 32l6.506-24.282a3 3 0 0 1 3.675-2.122v0a3 3 0 0 1 2.12 3.675L36.56 32"/><path fill="#fff" stroke="#fff" stroke-linecap="round" d="M36 32v6H13a3 3 0 1 1 0-6h23Z"/><path stroke="#fff" stroke-linecap="round" d="M21 44h8"/><rect width="10" height="20" x="6" y="4" fill="#fff" stroke="#fff" rx="5"/><path stroke="#000" stroke-linecap="round" d="M6 14h10"/><path stroke="#fff" stroke-linecap="round" d="M6 9v10M16 9v10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAirplaneWindowOne0)"/>`),
		g.Group(children),
	)
}