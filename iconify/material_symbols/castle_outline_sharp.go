package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CastleOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 21V9h2v2h2V3h2v2h2V3h2v2h2V3h2v2h2V3h2v8h2V9h2v12h-9v-5h-4v5H1Zm2-2h5v-5h8v5h5v-6h-4V7H7v6H3v6Zm6-7h2V9H9v3Zm4 0h2V9h-2v3Zm-1 1Z"/>`),
		g.Group(children),
	)
}