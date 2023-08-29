package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableRowHeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5h12a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2m0 4v3h5V9H3m7 0v3h5V9h-5m-7 5v3h5v-3H3m7 0v3h5v-3h-5m13 0V7h-4v2h2v3h-2v2h4Z"/>`),
		g.Group(children),
	)
}