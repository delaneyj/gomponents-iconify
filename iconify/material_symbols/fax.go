package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fax(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 20V4h10v5h1q1.25 0 2.125.875T22 12v8H8Zm-3.5 1q1.05 0 1.775-.725T7 18.5v-8q0-1.05-.725-1.775T4.5 8q-1.05 0-1.775.725T2 10.5v8q0 1.05.725 1.775T4.5 21ZM10 9h6V6h-6v3Zm6 5q.425 0 .713-.288T17 13q0-.425-.288-.713T16 12q-.425 0-.713.288T15 13q0 .425.288.713T16 14Zm3 0q.425 0 .713-.288T20 13q0-.425-.288-.713T19 12q-.425 0-.713.288T18 13q0 .425.288.713T19 14Zm-3 3q.425 0 .713-.288T17 16q0-.425-.288-.713T16 15q-.425 0-.713.288T15 16q0 .425.288.713T16 17Zm3 0q.425 0 .713-.288T20 16q0-.425-.288-.713T19 15q-.425 0-.713.288T18 16q0 .425.288.713T19 17Zm-9 0h4v-5h-4v5Z"/>`),
		g.Group(children),
	)
}