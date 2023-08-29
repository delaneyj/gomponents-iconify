package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTriangleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 12v10h6V12h4.5L12 2L4.5 12H9Zm-.5-2L12 5.333L15.5 10H13v10h-2V10H8.5Z"/>`),
		g.Group(children),
	)
}