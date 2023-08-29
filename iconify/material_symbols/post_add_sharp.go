package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PostAddSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 11V9h8v2H8Zm0 3v-2h8v2H8Zm0 3v-2h8v2H8Zm9-8V7h-2V5h2V3h2v2h2v2h-2v2h-2ZM3 21V3h11v2H5v14h14v-9h2v11H3Z"/>`),
		g.Group(children),
	)
}