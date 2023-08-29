package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2h18v16H6v2H4v-2h2v-2h14V4H4v18H2V2h2zm9 7h-2V7H9v2h2v2H9v2h2v-2h2v2h2v-2h-2V9zm0 0V7h2v2h-2z"/>`),
		g.Group(children),
	)
}