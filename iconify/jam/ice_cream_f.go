package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IceCreamF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 14v5a1 1 0 0 1-2 0v-5H3a3 3 0 0 1-3-3V5a5 5 0 1 1 10 0v6a3 3 0 0 1-3 3H6z"/>`),
		g.Group(children),
	)
}