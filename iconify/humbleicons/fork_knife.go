package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForkKnife(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 4v2a3 3 0 0 0 3 3m0 0V4m0 5v11M8 9a3 3 0 0 0 3-3V4m5 8V4c3 2 3 4 3 8h-3Zm0 0v8"/>`),
		g.Group(children),
	)
}