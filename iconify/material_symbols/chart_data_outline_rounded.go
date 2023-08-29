package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartDataOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.45 12.975l1.3 1.3q.275.275.7.275t.7-.275l2.85-2.85v.6q0 .425.288.7T17 13q.425 0 .713-.287T18 12V9q0-.425-.288-.713T17 8h-3.025q-.425 0-.7.288T13 9q0 .425.288.713T14 10h.575l-2.125 2.15l-1.3-1.3q-.275-.3-.7-.3t-.7.3L6.7 13.9q-.3.275-.3.7t.3.7q.275.3.7.3t.7-.3l2.35-2.325ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}