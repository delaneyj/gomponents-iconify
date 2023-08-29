package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutFooter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 5h20v14H2V5zm2 2v6h16V7H4zm16 8H4v2h16v-2z"/>`),
		g.Group(children),
	)
}