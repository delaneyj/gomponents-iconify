package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBarOverlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 28v-2h22V16H4v-2h14V4H4V2H2v26a2 2 0 0 0 2 2h26v-2Zm20-10v6H4v-2h16v-2H4v-2ZM16 6v6H4v-2h8V8H4V6Z"/>`),
		g.Group(children),
	)
}