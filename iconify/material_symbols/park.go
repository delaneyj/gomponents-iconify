package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Park(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.95 22h-3.9v-4H3l4-6H5l7-10l7 10h-2l4 6h-7.05v4Z"/>`),
		g.Group(children),
	)
}