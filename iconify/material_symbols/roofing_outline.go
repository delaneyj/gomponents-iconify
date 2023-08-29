package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoofingOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2 12l10-9l4 3.6V4h3v5.3l3 2.7h-3l-7-6.3L5 12H2Zm7 8v-6h6v6H9Zm2-2h2v-2h-2v2Zm1-1Z"/>`),
		g.Group(children),
	)
}