package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsFullscreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 0a1 1 0 0 1 1 1v6a1 1 0 0 1-2 0V3.414L11.414 10L18 16.586V13a1 1 0 0 1 2 0v6a1 1 0 0 1-1 1h-6a1 1 0 0 1 0-2h3.586L10 11.414L3.414 18H7a1 1 0 0 1 0 2H1a1 1 0 0 1-1-1v-6a1 1 0 0 1 2 0v3.586L8.586 10L2 3.414V7a1 1 0 1 1-2 0V1a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2H3.414L10 8.586L16.586 2H13a1 1 0 0 1 0-2h6z"/>`),
		g.Group(children),
	)
}