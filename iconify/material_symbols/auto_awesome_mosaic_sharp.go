package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoAwesomeMosaicSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21H3V3h8v18Zm2-10V3h8v8h-8Zm0 10v-8h8v8h-8Z"/>`),
		g.Group(children),
	)
}