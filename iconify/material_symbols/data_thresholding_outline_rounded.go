package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataThresholdingOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm12.275-2H19v-1.725L17.275 19ZM5.85 19h1.825l3-3H12.8l-3 3h1.6l3-3h2.125l-3 3h1.625l3-3H19V5H5v12.725L6.725 16H8.85l-3 3Zm4.825-8L8.35 13.3q-.275.275-.688.288T6.95 13.3q-.275-.275-.275-.7t.275-.7l3.025-3.025q.3-.3.7-.3t.7.3l1.3 1.3L15.65 7.2q.275-.275.688-.287t.712.287q.275.275.275.7t-.275.7l-3.675 3.7q-.3.3-.7.3t-.7-.3l-1.3-1.3ZM5 19V5v14Z"/>`),
		g.Group(children),
	)
}