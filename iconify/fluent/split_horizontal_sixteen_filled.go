package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitHorizontalSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.5 8a.5.5 0 0 0 0-1h-13a.5.5 0 0 0 0 1h13ZM11 1a2 2 0 0 1 2 2v3H3V3a2 2 0 0 1 2-2h6ZM3 9v3a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V9H3Z"/>`),
		g.Group(children),
	)
}