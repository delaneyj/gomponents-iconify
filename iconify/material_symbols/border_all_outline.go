package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderAllOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3Zm16-2v-6h-6v6h6Zm0-14h-6v6h6V5ZM5 5v6h6V5H5Zm0 14h6v-6H5v6Z"/>`),
		g.Group(children),
	)
}