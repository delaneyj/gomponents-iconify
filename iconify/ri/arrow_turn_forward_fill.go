package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnForwardFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 8v5a6 6 0 0 1-12 0V4H3v9a8 8 0 1 0 16 0V8h4l-5-6l-5 6h4Z"/>`),
		g.Group(children),
	)
}