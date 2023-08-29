package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnBackFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 16h-4l5 6l5-6h-4v-5a8 8 0 1 0-16 0v9h2v-9a6 6 0 1 1 12 0v5Z"/>`),
		g.Group(children),
	)
}