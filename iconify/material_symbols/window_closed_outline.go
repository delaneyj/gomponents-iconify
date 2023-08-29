package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowClosedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21V3h16v18H4Zm2-10h5v-1h2v1h5V5H6v6Zm0 8h12v-6H6v6Zm0 0h12H6Z"/>`),
		g.Group(children),
	)
}