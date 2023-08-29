package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StoreThreeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 13v7a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-7H2v-2l1-5h18l1 5v2h-1ZM5 13v6h14v-6H5Zm-.96-2h15.92l-.6-3H4.64l-.6 3ZM6 14h8v3H6v-3ZM3 3h18v2H3V3Z"/>`),
		g.Group(children),
	)
}