package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderTopVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 21h2v-2h-2m-4 2h2v-2h-2m4-2h2v-2h-2M3 5h18V3H3m16 10h2v-2h-2m0-2h2V7h-2M3 9h2V7H3m0 6h2v-2H3m0 10h2v-2H3m0-2h2v-2H3m8 6h2v-2h-2m-4 2h2v-2H7v2Z"/>`),
		g.Group(children),
	)
}