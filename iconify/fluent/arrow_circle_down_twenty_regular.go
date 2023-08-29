package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleDownTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m13.354 10.854l-3 3a.5.5 0 0 1-.708 0l-3-3a.5.5 0 0 1 .708-.708L9.5 12.293V6.5a.5.5 0 0 1 1 0v5.793l2.146-2.147a.5.5 0 0 1 .708.708ZM10 18a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm7-8a7 7 0 1 1-14 0a7 7 0 0 1 14 0Z"/>`),
		g.Group(children),
	)
}