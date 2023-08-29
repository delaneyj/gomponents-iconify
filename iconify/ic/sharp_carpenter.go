package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpCarpenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 1.5L3.11 5.39l8.13 11.67l-1.41 1.41l4.24 4.24l7.07-7.07L7 1.5zm5.66 16.97l4.24-4.24l1.41 1.41l-4.24 4.24l-1.41-1.41z"/>`),
		g.Group(children),
	)
}