package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSmallFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 21h-4v-2h4v-2h-4v-6h6v2h-4v2h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}