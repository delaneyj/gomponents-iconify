package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 17v2h-8v-2h8m-11-1L2 9l9-7l9 7l-9 7m0 2.54l1-.79V18c0 .71.12 1.39.35 2L11 21.07l-9-7l1.62-1.26L11 18.54Z"/>`),
		g.Group(children),
	)
}