package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Key(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 11c0 .732.166 1.424.449 2.051L5 17v1.5S5.896 20 7 20h2v-2h2v-2h2.5c2.762 0 5-2.238 5-5s-2.238-5-5-5s-5 2.238-5 5zm5 2a2 2 0 1 1 .001-4.001A2 2 0 0 1 13.5 13z"/>`),
		g.Group(children),
	)
}