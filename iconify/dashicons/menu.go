package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Menu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 15h14v-2H3v2zM3 5v2h14V5H3zm0 6h14V9H3v2z"/>`),
		g.Group(children),
	)
}