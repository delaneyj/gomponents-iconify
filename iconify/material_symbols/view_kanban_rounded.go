package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewKanbanRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 17q.425 0 .713-.288T9 16V8q0-.425-.288-.713T8 7q-.425 0-.713.288T7 8v8q0 .425.288.713T8 17Zm4-5q.425 0 .713-.288T13 11V8q0-.425-.288-.713T12 7q-.425 0-.713.288T11 8v3q0 .425.288.713T12 12Zm4 3q.425 0 .713-.288T17 14V8q0-.425-.288-.713T16 7q-.425 0-.713.288T15 8v6q0 .425.288.713T16 15ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}