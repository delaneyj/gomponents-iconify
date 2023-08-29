package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBagThreeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.505 2h11a1 1 0 0 1 .8.4l2.7 3.6v15a1 1 0 0 1-1 1h-16a1 1 0 0 1-1-1V6l2.7-3.6a1 1 0 0 1 .8-.4Zm12 4l-1.5-2h-10l-1.5 2h13Zm-9.5 4h-2v2a5 5 0 0 0 10 0v-2h-2v2a3 3 0 0 1-6 0v-2Z"/>`),
		g.Group(children),
	)
}