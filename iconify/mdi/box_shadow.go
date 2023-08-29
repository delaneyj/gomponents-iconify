package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxShadow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h15v15H3V3m16 16h2v2h-2v-2m0-3h2v2h-2v-2m0-3h2v2h-2v-2m0-3h2v2h-2v-2m0-3h2v2h-2V7m-3 12h2v2h-2v-2m-3 0h2v2h-2v-2m-3 0h2v2h-2v-2m-3 0h2v2H7v-2Z"/>`),
		g.Group(children),
	)
}