package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shipping(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 8h-2V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v10h2a3 3 0 1 0 6 0h4a3 3 0 1 0 6 0h2v-5l-4-4zM7 18.5a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 7 18.5zM4 14V7h10v7H4zm13 4.5a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 17 18.5z"/>`),
		g.Group(children),
	)
}