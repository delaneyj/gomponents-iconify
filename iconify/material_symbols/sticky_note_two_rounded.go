package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickyNoteTwoRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14 19l5-5h-4q-.425 0-.713.288T14 15v4Zm-9 2q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v9.175q0 .4-.15.763t-.425.637l-4.85 4.85q-.275.275-.637.425t-.763.15H5Zm3-7h3q.425 0 .713-.288T12 13q0-.425-.288-.713T11 12H8q-.425 0-.713.288T7 13q0 .425.288.713T8 14Zm0-4h8q.425 0 .713-.288T17 9q0-.425-.288-.713T16 8H8q-.425 0-.713.288T7 9q0 .425.288.713T8 10Z"/>`),
		g.Group(children),
	)
}