package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RememberMe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 23q-.825 0-1.413-.588T5 21V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v18q0 .825-.588 1.413T17 23H7Zm5-9q1.35 0 2.613.313T17 15.2V6H7v9.2q1.125-.575 2.388-.887T12 14Zm0-1q-1.25 0-2.125-.875T9 10q0-1.25.875-2.125T12 7q1.25 0 2.125.875T15 10q0 1.25-.875 2.125T12 13Z"/>`),
		g.Group(children),
	)
}