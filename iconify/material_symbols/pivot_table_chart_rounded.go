package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PivotTableChartRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 8V3h9q.825 0 1.413.588T21 5v3H10ZM5 21q-.825 0-1.413-.588T3 19v-9h5v11H5ZM3 8V5q0-.825.588-1.413T5 3h3v5H3Zm9.3 13.3l-2.6-2.6q-.15-.15-.212-.325T9.425 18q0-.2.063-.375T9.7 17.3l2.625-2.625Q12.6 14.4 13 14.4t.7.3q.275.275.275.688t-.275.712l-.85.9H15q.825 0 1.413-.588T17 15v-2.2l-.925.925Q15.8 14 15.4 14t-.7-.3q-.275-.275-.275-.7t.275-.7l2.6-2.6q.15-.15.325-.212T18 9.425q.2 0 .375.063t.325.212l2.625 2.625q.275.275.275.675t-.3.7q-.275.275-.7.275t-.7-.275l-.9-.9V15q0 1.65-1.175 2.825T15 19h-2.15l.875.9q.275.3.275.7t-.3.7q-.275.275-.7.275t-.7-.275Z"/>`),
		g.Group(children),
	)
}