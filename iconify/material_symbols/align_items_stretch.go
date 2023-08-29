package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignItemsStretch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 18V6h3v12H7Zm7 0V6h3v12h-3ZM2 4V2h20v2H2Zm0 18v-2h20v2H2Z"/>`),
		g.Group(children),
	)
}