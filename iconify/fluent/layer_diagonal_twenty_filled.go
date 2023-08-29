package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayerDiagonalTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.447 2.107a1.5 1.5 0 0 1 2.05 1.301L5.42 5.814A2.25 2.25 0 0 0 4 7.906v6.006a1.5 1.5 0 0 1-2-1.414V6.755a2.5 2.5 0 0 1 1.58-2.324l5.867-2.324Zm3 2a1.5 1.5 0 0 1 2.052 1.395v.102L9.264 7.678A2 2 0 0 0 8 9.538v5.98l-.947.375A1.5 1.5 0 0 1 5 14.498V8.416a2 2 0 0 1 1.264-1.86l6.183-2.449ZM18 7.501v6.623a1.5 1.5 0 0 1-.948 1.395l-6 2.376A1.5 1.5 0 0 1 9 16.5V9.877a1.5 1.5 0 0 1 .948-1.394l6-2.377A1.5 1.5 0 0 1 18 7.501Z"/>`),
		g.Group(children),
	)
}