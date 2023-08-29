package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowMaximizeSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.477 2a.75.75 0 0 0 0 1.5h2.962L3.5 11.44V8.476a.75.75 0 0 0-1.5 0v4.673c0 .47.38.85.85.85h4.673a.75.75 0 0 0 0-1.5H4.56L12.5 4.56v2.963a.75.75 0 0 0 1.5 0V2.85a.85.85 0 0 0-.85-.85H8.477Z"/>`),
		g.Group(children),
	)
}