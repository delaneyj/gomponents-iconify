package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceGeometricTriangleGeometricTriangleShapeDesignShapes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.46 2a.55.55 0 0 0-.92 0l-6 9.5a.5.5 0 0 0 0 .5a.54.54 0 0 0 .46.25h12a.54.54 0 0 0 .46-.25a.5.5 0 0 0 0-.5Z"/>`),
		g.Group(children),
	)
}