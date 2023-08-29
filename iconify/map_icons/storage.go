package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Storage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M3 6h44v5H3zm3 7v33h38V13H6zm26 9H17v-3h15v3z"/>`),
		g.Group(children),
	)
}