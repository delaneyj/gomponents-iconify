package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickyNoteTwoOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19h9v-4q0-.425.288-.713T15 14h4V5H5v14Zm0 2q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v9.175q0 .4-.15.763t-.425.637l-4.85 4.85q-.275.275-.637.425t-.763.15H5Zm6-7H8q-.425 0-.713-.288T7 13q0-.425.288-.713T8 12h3q.425 0 .713.288T12 13q0 .425-.288.713T11 14Zm5-4H8q-.425 0-.713-.288T7 9q0-.425.288-.713T8 8h8q.425 0 .713.288T17 9q0 .425-.288.713T16 10ZM5 19V5v14Z"/>`),
		g.Group(children),
	)
}