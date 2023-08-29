package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleLeftSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m7.146 10.854l-2.5-2.5a.5.5 0 0 1 0-.708l2.5-2.5a.5.5 0 1 1 .708.708L6.207 7.5H11a.5.5 0 0 1 0 1H6.207l1.647 1.646a.5.5 0 0 1-.708.708ZM1 8a7 7 0 1 0 14 0A7 7 0 0 0 1 8Zm7 6A6 6 0 1 1 8 2a6 6 0 0 1 0 12Z"/>`),
		g.Group(children),
	)
}