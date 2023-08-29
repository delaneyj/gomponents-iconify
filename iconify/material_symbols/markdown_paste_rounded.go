package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkdownPasteRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 21q-.425 0-.713-.288T12 20v-6q0-.825.588-1.413T14 12h6q.825 0 1.413.588T22 14v6q0 .425-.288.713T21 21q-.425 0-.713-.288T20 20v-6h-2v4q0 .425-.288.713T17 19q-.425 0-.713-.288T16 18v-4h-2v6q0 .425-.288.713T13 21Zm-8 0q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h4.175q.275-.875 1.075-1.438T12 1q1 0 1.788.563T14.85 3H19q.825 0 1.413.588T21 5v4q0 .425-.288.713T20 10q-.425 0-.713-.288T19 9V5h-2v2q0 .425-.288.713T16 8H8q-.425 0-.713-.288T7 7V5H5v14h4q.425 0 .713.288T10 20q0 .425-.288.713T9 21H5Zm7-16q.425 0 .713-.288T13 4q0-.425-.288-.713T12 3q-.425 0-.713.288T11 4q0 .425.288.713T12 5Z"/>`),
		g.Group(children),
	)
}