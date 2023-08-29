package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderRightVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 5h2V3h-2m4 2h2V3h-2m0 18h2v-2h-2m4 2h2V3h-2M3 9h2V7H3m0 10h2v-2H3m0-2h2v-2H3m8 10h2v-2h-2m-8 2h2v-2H3M7 5h2V3H7M3 5h2V3H3m4 18h2v-2H7v2Z"/>`),
		g.Group(children),
	)
}