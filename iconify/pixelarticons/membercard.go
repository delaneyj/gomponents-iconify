package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Membercard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 3h20v14h-7v3h-2v-2h-2v2H9v-3H2V3zm2 2v4h16V5H4zm16 8H4v2h16v-2z"/>`),
		g.Group(children),
	)
}