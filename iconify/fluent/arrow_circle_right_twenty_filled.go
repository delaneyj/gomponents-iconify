package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleRightTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 18a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm.854-4.646a.5.5 0 0 1-.708-.708l2.147-2.146H6.5a.5.5 0 0 1 0-1h5.793l-2.147-2.146a.5.5 0 0 1 .708-.708l3 3a.5.5 0 0 1 0 .708l-3 3Z"/>`),
		g.Group(children),
	)
}