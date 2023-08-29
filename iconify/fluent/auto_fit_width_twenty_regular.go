package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoFitWidthTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 4a.5.5 0 0 0-1 0v11a.5.5 0 0 0 1 0V4ZM3 4a.5.5 0 0 0-1 0v11a.5.5 0 0 0 1 0V4Zm11.854 5.146a.5.5 0 0 1 0 .708l-2 2a.5.5 0 0 1-.708-.708L13.293 10H6.707l1.147 1.146a.5.5 0 0 1-.708.708l-2-2a.5.5 0 0 1 0-.708l2-2a.5.5 0 1 1 .708.708L6.707 9h6.586l-1.147-1.146a.5.5 0 0 1 .708-.708l2 2Z"/>`),
		g.Group(children),
	)
}