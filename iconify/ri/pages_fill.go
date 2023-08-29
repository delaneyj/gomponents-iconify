package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PagesFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 22H4a1 1 0 0 1-1-1V8h18v13a1 1 0 0 1-1 1Zm1-16H3V3a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v3ZM7 11v4h4v-4H7Zm0 6v2h10v-2H7Zm6-5v2h4v-2h-4Z"/>`),
		g.Group(children),
	)
}