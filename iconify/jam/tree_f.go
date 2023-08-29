package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 13.758v6.1a1 1 0 0 1-2 0v-6.1A5 5 0 0 1 2.118 4.77a5.002 5.002 0 0 1 9.764 0A5 5 0 0 1 8 13.758z"/>`),
		g.Group(children),
	)
}