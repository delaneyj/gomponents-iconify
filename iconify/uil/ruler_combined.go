package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerCombined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V10h11a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1Zm-1 6h-2V7a1 1 0 0 0-2 0v1h-2V7a1 1 0 0 0-2 0v1h-2V7a1 1 0 0 0-2 0v1H7a1 1 0 0 0 0 2h1v2H7a1 1 0 0 0 0 2h1v2H7a1 1 0 0 0 0 2h1v2H4V4h16Z"/>`),
		g.Group(children),
	)
}