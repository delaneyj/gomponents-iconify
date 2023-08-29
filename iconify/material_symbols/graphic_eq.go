package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraphicEq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 18V6h2v12H7Zm4 4V2h2v20h-2Zm-8-8v-4h2v4H3Zm12 4V6h2v12h-2Zm4-4v-4h2v4h-2Z"/>`),
		g.Group(children),
	)
}