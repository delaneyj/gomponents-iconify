package mi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Layout(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H3.987A2 2 0 0 1 2 19V5zm2 4h16V5H4v4zm4 2H4v8h4v-8zm2 8h10v-8H10v8z"/>`),
		g.Group(children),
	)
}