package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTriangleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 12V2h6v10h4.5L12 22L4.5 12H9Zm-.5 2l3.5 4.667L15.5 14H13V4h-2v10H8.5Z"/>`),
		g.Group(children),
	)
}