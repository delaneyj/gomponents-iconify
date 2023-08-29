package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WbShade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 20L14 14.5V12l8 8h-2.5ZM14 20v-3l3 3h-3ZM4 20V10H2l6-6l6 6h-2v10H4Zm3-6h2v-4H7v4Z"/>`),
		g.Group(children),
	)
}