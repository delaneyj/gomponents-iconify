package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListCheckmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 15.5L5 20l-2.5-2.5l1.061-1.061L5 17.879l3.439-3.439L9.5 15.5zM10 5v2h11V5H10zm0 14h11v-2H10v2zm0-6h11v-2H10v2zM8.439 8.439L5 11.879L3.561 10.44L2.5 11.5L5 14l4.5-4.5l-1.061-1.061zm0-6L5 5.879l-1.439-1.44L2.5 5.5L5 8l4.5-4.5l-1.061-1.061z"/>`),
		g.Group(children),
	)
}