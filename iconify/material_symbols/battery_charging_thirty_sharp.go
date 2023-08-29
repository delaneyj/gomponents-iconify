package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryChargingThirtySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 22v-3H14l3.5-5v3H20l-3.5 5ZM7 22V4h3V2h4v2h3v8q-.525 0-1.025.088T15 12.35V6H9v10h2.35q-.175.475-.263.975T11 18q0 1.15.4 2.175T12.525 22H7Z"/>`),
		g.Group(children),
	)
}