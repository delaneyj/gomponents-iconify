package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickyNoteOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v10.175q0 .4-.15.763t-.425.637l-3.85 3.85q-.275.275-.638.425t-.762.15H5Zm10-2v-2q0-.825.588-1.413T17 15h2V5H5v14h10Zm-4-9v5q0 .425.288.713T12 16q.425 0 .713-.288T13 15v-5h2q.425 0 .713-.288T16 9q0-.425-.288-.713T15 8H9q-.425 0-.713.288T8 9q0 .425.288.713T9 10h2Zm4 9ZM5 19V5v14Z"/>`),
		g.Group(children),
	)
}