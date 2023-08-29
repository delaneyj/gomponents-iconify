package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaletteAdvanced(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 22H10v-2h12v2M2 22v-2h7v2H2m16-4v-8h4v8h-4m0-15h4v6h-4V3M2 18V3h14v15H2m7-3.44a3 3 0 0 0 3-3c0-2-3-5.37-3-5.37s-3 3.37-3 5.37a3 3 0 0 0 3 3Z"/>`),
		g.Group(children),
	)
}