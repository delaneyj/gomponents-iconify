package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PivotTableChartSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 8V3h11v5H10ZM3 21V10h5v11H3ZM3 8V3h5v5H3Zm10 14l-4-4l4-4l1.4 1.4l-1.55 1.6H15q.825 0 1.413-.588T17 15v-2.2l-1.6 1.6L14 13l4-4l4 4l-1.4 1.4l-1.6-1.6V15q0 1.65-1.175 2.825T15 19h-2.15l1.55 1.6L13 22Z"/>`),
		g.Group(children),
	)
}