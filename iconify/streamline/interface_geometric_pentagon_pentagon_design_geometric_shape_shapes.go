package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceGeometricPentagonPentagonDesignGeometricShapeShapes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 6L7 .5L.5 6l3 7.5h7l3-7.5z"/>`),
		g.Group(children),
	)
}