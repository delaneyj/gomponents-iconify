package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableRows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 21H3v-4.65h18V21Zm0-6.65H3V9.625h18v4.725Zm0-6.725H3V3h18v4.625Z"/>`),
		g.Group(children),
	)
}