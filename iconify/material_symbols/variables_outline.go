package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VariablesOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 17v-3h-3v-2h3V9h2v3h3v2h-3v3h-2ZM3 14V4h18v3h-2V6H5v6h9v2H3Zm2-4V6v6v-2Z"/>`),
		g.Group(children),
	)
}