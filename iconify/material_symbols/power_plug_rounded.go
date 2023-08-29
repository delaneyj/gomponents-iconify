package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerPlugRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.5 18l-2.925-2.925q-.275-.275-.425-.637T6 13.675V9q0-.825.588-1.413T8 7h1L8 8V4q0-.425.288-.713T9 3q.425 0 .713.288T10 4v3h4V4q0-.425.288-.713T15 3q.425 0 .713.288T16 4v4l-1-1h1q.825 0 1.413.588T18 9v4.675q0 .4-.15.763t-.425.637L14.5 18v2q0 .425-.288.713T13.5 21h-3q-.425 0-.713-.288T9.5 20v-2Z"/>`),
		g.Group(children),
	)
}