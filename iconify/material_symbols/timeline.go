package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timeline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 18q-.825 0-1.413-.588T1 16q0-.825.588-1.413T3 14h.263q.112 0 .237.05L8.05 9.5Q8 9.375 8 9.262V9q0-.825.588-1.413T10 7q.825 0 1.413.588T12 9q0 .05-.05.5l2.55 2.55q.125-.05.238-.05h.525q.112 0 .237.05l3.55-3.55Q19 8.375 19 8.262V8q0-.825.588-1.413T21 6q.825 0 1.413.588T23 8q0 .825-.588 1.413T21 10h-.263q-.112 0-.237-.05l-3.55 3.55q.05.125.05.238V14q0 .825-.588 1.413T15 16q-.825 0-1.413-.588T13 14v-.263q0-.112.05-.237l-2.55-2.55q-.125.05-.238.05H10q-.05 0-.5-.05L4.95 15.5q.05.125.05.238V16q0 .825-.588 1.413T3 18Z"/>`),
		g.Group(children),
	)
}