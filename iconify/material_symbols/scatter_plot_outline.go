package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScatterPlotOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 21q-1.65 0-2.825-1.175T13 17q0-1.65 1.175-2.825T17 13q1.65 0 2.825 1.175T21 17q0 1.65-1.175 2.825T17 21Zm0-2q.825 0 1.413-.588T19 17q0-.825-.588-1.413T17 15q-.825 0-1.413.588T15 17q0 .825.588 1.413T17 19ZM7 18q-1.65 0-2.825-1.175T3 14q0-1.65 1.175-2.825T7 10q1.65 0 2.825 1.175T11 14q0 1.65-1.175 2.825T7 18Zm0-2q.825 0 1.413-.588T9 14q0-.825-.588-1.413T7 12q-.825 0-1.413.588T5 14q0 .825.588 1.413T7 16Zm4-6q-1.65 0-2.825-1.175T7 6q0-1.65 1.175-2.825T11 2q1.65 0 2.825 1.175T15 6q0 1.65-1.175 2.825T11 10Zm0-2q.825 0 1.413-.588T13 6q0-.825-.588-1.413T11 4q-.825 0-1.413.588T9 6q0 .825.588 1.413T11 8Zm6 9ZM7 14Zm4-8Z"/>`),
		g.Group(children),
	)
}