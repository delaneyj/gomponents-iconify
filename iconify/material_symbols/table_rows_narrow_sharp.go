package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableRowsNarrowSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 15.5h18V13H3v2.5ZM3 11h18V8.5H3V11Zm0-4.5h18V4H3v2.5ZM21 21V3v18ZM3 20h18v-2.5H3V20Z"/>`),
		g.Group(children),
	)
}