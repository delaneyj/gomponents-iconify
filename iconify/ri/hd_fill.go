package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HdFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1Zm4.5 8.25V9H6v6h1.5v-2.25h2V15H11V9H9.5v2.25h-2Zm7-.75H16a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5h-1.5v-3ZM13 9v6h3a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2h-3Z"/>`),
		g.Group(children),
	)
}