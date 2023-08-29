package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartMinusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 11.475ZM11 21l-3.175-2.85q-1.8-1.625-3.088-2.9t-2.125-2.4q-.837-1.125-1.225-2.175T1 8.475q0-2.35 1.575-3.913T6.5 3q1.3 0 2.475.55T11 5.1q.85-1 2.025-1.55T15.5 3q2.1 0 3.825 1.475t1.725 4q0 .35-.05.738t-.15.787h-2.125q.125-.45.2-.85T19 8.4q0-1.875-1.25-2.638T15.5 5q-1.275 0-2.2.688T11.575 7.5h-1.15Q9.65 6.375 8.662 5.687T6.5 5q-1.425 0-2.463.988T3 8.474q0 .825.35 1.675t1.25 1.963q.9 1.112 2.45 2.6T11 18.3q.65-.575 1.525-1.325t1.4-1.25l.225.225l.488.488l.487.487l.225.225q-.55.5-1.4 1.238t-1.5 1.312L11 21Zm4-7v-2h8v2h-8Z"/>`),
		g.Group(children),
	)
}