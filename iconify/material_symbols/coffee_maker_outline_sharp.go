package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoffeeMakerOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22V2h16v2h-2v3H8V4H6v16h4.05q-.95-.675-1.5-1.713T8 16v-5h10v5q0 1.25-.55 2.288T15.95 20H20v2H4Zm9-3q1.25 0 2.125-.875T16 16v-3h-6v3q0 1.25.875 2.125T13 19Zm0-9q.425 0 .713-.288T14 9q0-.425-.288-.713T13 8q-.425 0-.713.288T12 9q0 .425.288.713T13 10Zm0 3Z"/>`),
		g.Group(children),
	)
}