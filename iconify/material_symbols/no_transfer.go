package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoTransfer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.475 23.3L16.3 19.125h2.85V20q0 .425-.287.713T18.15 21H17q-.425 0-.712-.288T16 20v-1H8v1q0 .425-.288.713T7 21H6q-.425 0-.713-.288T5 20v-2.05q-.45-.5-.725-1.113T4 15.5V6.825L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425Zm-.75-6.45L12.875 10H18V7H9.875l-4-4q.975-.5 2.488-.75T12 2q4.3 0 6.15.925T20 6v9.5q0 .35-.075.688t-.2.662ZM8.5 16q.625 0 1.063-.438T10 14.5q0-.625-.438-1.063T8.5 13q-.625 0-1.063.438T7 14.5q0 .625.438 1.063T8.5 16ZM6 10h1.175L6 8.825V10Z"/>`),
		g.Group(children),
	)
}