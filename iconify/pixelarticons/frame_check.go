package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrameCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 3h20v18H2V3zm18 16V7H4v12h16zm-4-9h-2v2h-2v2h-2v-2H8v2h2v2h2v-2h2v-2h2v-2z"/>`),
		g.Group(children),
	)
}