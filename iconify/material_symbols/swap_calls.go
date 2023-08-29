package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapCalls(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 19l-4-4l1.4-1.45l1.6 1.6V8q0-1.65 1.175-2.825T9 4q1.65 0 2.825 1.175T13 8v7q0 .825.588 1.413T15 17q.825 0 1.413-.588T17 15V7.85l-1.6 1.6L14 8l4-4l4 4l-1.4 1.45l-1.6-1.6V15q0 1.65-1.175 2.825T15 19q-1.65 0-2.825-1.175T11 15V8q0-.825-.588-1.413T9 6q-.825 0-1.413.588T7 8v7.15l1.6-1.6L10 15l-4 4Z"/>`),
		g.Group(children),
	)
}