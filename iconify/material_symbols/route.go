package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Route(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 21q-1.65 0-2.825-1.175T5 17V8.825Q4.125 8.5 3.562 7.738T3 6q0-1.25.875-2.125T6 3q1.25 0 2.125.875T9 6q0 .975-.563 1.738T7 8.825V17q0 .825.588 1.413T9 19q.825 0 1.413-.588T11 17V7q0-1.65 1.175-2.825T15 3q1.65 0 2.825 1.175T19 7v8.175q.875.325 1.438 1.088T21 18q0 1.25-.875 2.125T18 21q-1.25 0-2.125-.875T15 18q0-.975.563-1.75T17 15.175V7q0-.825-.588-1.413T15 5q-.825 0-1.413.588T13 7v10q0 1.65-1.175 2.825T9 21Z"/>`),
		g.Group(children),
	)
}