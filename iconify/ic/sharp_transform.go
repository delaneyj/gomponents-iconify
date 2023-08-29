package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpTransform(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 18v-2H8V4h2L7 1L4 4h2v2H2v2h4v10h10v2h-2l3 3l3-3h-2v-2h4zM10 8h6v6h2V6h-8v2z"/>`),
		g.Group(children),
	)
}