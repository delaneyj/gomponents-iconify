package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataTable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 8h18V4H3v4Zm0 6h18v-4H3v4Zm0 6h18v-4H3v4ZM4 7V5h2v2H4Zm0 6v-2h2v2H4Zm0 6v-2h2v2H4Z"/>`),
		g.Group(children),
	)
}