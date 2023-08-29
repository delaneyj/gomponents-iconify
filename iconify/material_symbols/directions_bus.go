package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsBus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21q-.425 0-.713-.288T5 20v-2.05q-.45-.5-.725-1.113T4 15.5V6q0-2.075 1.925-3.038T12 2q4.3 0 6.15.925T20 6v9.5q0 .725-.275 1.338T19 17.95V20q0 .425-.288.713T18 21h-1q-.425 0-.713-.288T16 20v-1H8v1q0 .425-.288.713T7 21H6Zm0-11h12V7H6v3Zm2.5 6q.625 0 1.063-.438T10 14.5q0-.625-.438-1.063T8.5 13q-.625 0-1.063.438T7 14.5q0 .625.438 1.063T8.5 16Zm7 0q.625 0 1.063-.438T17 14.5q0-.625-.438-1.063T15.5 13q-.625 0-1.063.438T14 14.5q0 .625.438 1.063T15.5 16Z"/>`),
		g.Group(children),
	)
}