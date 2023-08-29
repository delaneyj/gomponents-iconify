package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicationLiquidSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5V3h12v2H3Zm4.5 12.5h3V15H13v-3h-2.5V9.5h-3V12H5v3h2.5v2.5ZM2 21V6h14v15H2Zm17 0v-7.25q-.875-.425-1.438-1.413T17 10q0-1.7.863-2.85T20 6q1.275 0 2.138 1.15T23 10q0 1.35-.563 2.337T21 13.75V21h-2Z"/>`),
		g.Group(children),
	)
}