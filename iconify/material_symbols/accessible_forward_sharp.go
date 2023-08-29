package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccessibleForwardSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22q-2.075 0-3.538-1.463T3 17q0-2.075 1.463-3.538T8 12v2q-1.25 0-2.125.875T5 17q0 1.25.875 2.125T8 20q1.25 0 2.125-.875T11 17h2q0 2.075-1.463 3.538T8 22Zm9-1v-5h-6q-1.1 0-1.7-.938T9.15 13.1L11 9H8.725l-.6 1.55L6.2 10l1.175-3h6.575q1.125 0 1.713.913T15.8 9.85l-1.65 3.65H18l1 1V21h-2ZM16 6.5q-.825 0-1.413-.588T14 4.5q0-.825.588-1.413T16 2.5q.825 0 1.413.588T18 4.5q0 .825-.588 1.413T16 6.5Z"/>`),
		g.Group(children),
	)
}