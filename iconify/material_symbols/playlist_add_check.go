package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistAddCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 16v-2h8v2H3Zm0-4v-2h12v2H3Zm0-4V6h12v2H3Zm13.35 11l-3.55-3.55l1.4-1.4l2.15 2.1l4.25-4.25l1.4 1.45L16.35 19Z"/>`),
		g.Group(children),
	)
}