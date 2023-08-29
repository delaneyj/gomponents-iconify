package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OvenGen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 12v7q0 .825.588 1.413T5 21h14q.825 0 1.413-.588T21 19v-7h-4v5H7v-5H3Zm6 3h6v-3H9v3Zm-6-5h18V5q0-.825-.588-1.413T19 3H5q-.825 0-1.413.588T3 5v5Zm5-2q-.425 0-.713-.288T7 7q0-.425.288-.713T8 6q.425 0 .713.288T9 7q0 .425-.288.713T8 8Zm4 0q-.425 0-.713-.288T11 7q0-.425.288-.713T12 6q.425 0 .713.288T13 7q0 .425-.288.713T12 8Zm4 0q-.425 0-.713-.288T15 7q0-.425.288-.713T16 6q.425 0 .713.288T17 7q0 .425-.288.713T16 8Z"/>`),
		g.Group(children),
	)
}