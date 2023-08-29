package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeliveryFastFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 3a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v5h4a5 5 0 0 1 5 5v4a3.001 3.001 0 0 1-2.129 2.872a3 3 0 0 1-5.7.128H8.83a3 3 0 0 1-5.7-.128A3.001 3.001 0 0 1 1 17v-4h6a1 1 0 1 0 0-2H1V9h4a1 1 0 0 0 0-2H1V3Zm13 15h1.171a3 3 0 0 1 5.536-.293A.997.997 0 0 0 21 17v-4a3 3 0 0 0-3-3h-4v8Zm-7 1a1 1 0 1 0-2 0a1 1 0 0 0 2 0Zm10.293-.707A.994.994 0 0 0 17 19a1 1 0 1 0 .293-.707Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}