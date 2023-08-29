package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3 7v2h26V7H3zm0 4v2h18v-2H3zm0 4v2h26v-2H3zm0 4v2h18v-2H3zm0 4v2h26v-2H3z"/>`),
		g.Group(children),
	)
}