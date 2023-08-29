package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 16v-2h7v2H3Zm0-4v-2h11v2H3Zm0-4V6h11v2H3Zm13 12v-4h-4v-2h4v-4h2v4h4v2h-4v4h-2Z"/>`),
		g.Group(children),
	)
}