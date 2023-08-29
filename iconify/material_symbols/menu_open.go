package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 18v-2h13v2H3Zm16.6-1l-5-5l5-5L21 8.4L17.4 12l3.6 3.6l-1.4 1.4ZM3 13v-2h10v2H3Zm0-5V6h13v2H3Z"/>`),
		g.Group(children),
	)
}