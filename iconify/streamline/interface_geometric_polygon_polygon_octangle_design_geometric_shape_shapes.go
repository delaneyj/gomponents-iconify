package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceGeometricPolygonPolygonOctangleDesignGeometricShapeShapes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 13.5h-5l-4-4v-5l4-4h5l4 4v5l-4 4z"/>`),
		g.Group(children),
	)
}