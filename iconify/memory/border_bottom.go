package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M20 12h-2v-2h2v2M4 12H2v-2h2v2m8-8h-2V2h2v2m4 0h-2V2h2v2m4 4h-2V6h2v2m0-4h-2V2h2v2m0 12h-2v-2h2v2M4 16H2v-2h2v2M8 4H6V2h2v2M4 4H2V2h2v2m0 4H2V6h2v2m16 12H2v-2h18v2Z"/>`),
		g.Group(children),
	)
}