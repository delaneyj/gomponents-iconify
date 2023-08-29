package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistPlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 16v-2h8v2H3Zm0-4v-2h12v2H3Zm0-4V6h12v2H3Zm13 13v-8l6 4l-6 4Z"/>`),
		g.Group(children),
	)
}