package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignNone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 5h14V3H3v2zm10 8V7H3v6h10zM3 17h14v-2H3v2z"/>`),
		g.Group(children),
	)
}