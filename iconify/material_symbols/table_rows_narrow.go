package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableRowsNarrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 15.5V13h18v2.5H3ZM3 11V8.5h18V11H3Zm0-4.5V4h18v2.5H3ZM3 20v-2.5h18V20H3Z"/>`),
		g.Group(children),
	)
}