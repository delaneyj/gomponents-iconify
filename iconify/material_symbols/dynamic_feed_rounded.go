package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DynamicFeedRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.825 0-1.413-.588T2 19v-6q0-.425.288-.713T3 12q.425 0 .713.288T4 13v6h8q.425 0 .713.288T13 20q0 .425-.288.713T12 21H4Zm4-4q-.825 0-1.413-.588T6 15V9q0-.425.288-.713T7 8q.425 0 .713.288T8 9v6h8q.425 0 .713.288T17 16q0 .425-.288.713T16 17H8Zm4-4q-.825 0-1.413-.588T10 11V5q0-.825.588-1.413T12 3h8q.825 0 1.413.588T22 5v6q0 .825-.588 1.413T20 13h-8Zm0-2h8V7h-8v4Z"/>`),
		g.Group(children),
	)
}