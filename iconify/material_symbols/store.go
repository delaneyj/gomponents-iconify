package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Store(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 6V4h16v2H4Zm0 14v-6H3v-2l1-5h16l1 5v2h-1v6h-2v-6h-4v6H4Zm2-2h6v-4H6v4Z"/>`),
		g.Group(children),
	)
}