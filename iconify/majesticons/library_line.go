package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LibraryLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 9v11M4 9h16M4 9H3l9-5l9 5h-1M4 20h16M4 20H3m17 0V9m0 11h1M8 13v3m4 0v-3m4 0v3"/>`),
		g.Group(children),
	)
}