package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermostatCarbonOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22.8L1.2 12L12 1.2L22.8 12L12 22.8Zm0-2.8l8-8l-8-8l-8 8l8 8Zm0-8Z"/>`),
		g.Group(children),
	)
}