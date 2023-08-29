package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrameFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 7v10H7V7h10ZM5 7v10H3a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h10v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2V7h2a1 1 0 1 0 0-2h-2V3a1 1 0 1 0-2 0v2H7V3a1 1 0 0 0-2 0v2H3a1 1 0 0 0 0 2h2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}