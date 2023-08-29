package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutHeader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 19h20V5H2v14zm2-2v-6h16v6H4zm16-8H4V7h16v2z"/>`),
		g.Group(children),
	)
}