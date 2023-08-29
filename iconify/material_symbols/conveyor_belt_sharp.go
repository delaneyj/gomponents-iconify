package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConveyorBeltSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-1.25 0-2.125-.875T2 18q0-1.25.875-2.125T5 15h14q1.25 0 2.125.875T22 18q0 1.25-.875 2.125T19 21H5Zm0-2h14q.425 0 .713-.288T20 18q0-.425-.288-.713T19 17H5q-.425 0-.713.288T4 18q0 .425.288.713T5 19Zm4-6V3h10v10H9Zm-7-2.05V9.025h4.925v1.925H2ZM12 8h4V6.025h-4V8ZM4 8h2.925V6.025H4V8Z"/>`),
		g.Group(children),
	)
}