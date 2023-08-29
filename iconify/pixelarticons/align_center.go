package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 5H4v2h16V5zm-4 4H8v2h8V9zM4 13h16v2H4v-2zm12 4H8v2h8v-2z"/>`),
		g.Group(children),
	)
}