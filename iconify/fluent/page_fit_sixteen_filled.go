package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageFitSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V5Zm2 1v1a.5.5 0 0 0 1 0V6h1a.5.5 0 0 0 0-1H4a1 1 0 0 0-1 1Zm9-1h-1a.5.5 0 0 0 0 1h1v1a.5.5 0 0 0 1 0V6a1 1 0 0 0-1-1Zm0 6a1 1 0 0 0 1-1V9a.5.5 0 0 0-1 0v1h-1a.5.5 0 0 0 0 1h1Zm-8 0h1a.5.5 0 0 0 0-1H4V9a.5.5 0 0 0-1 0v1a1 1 0 0 0 1 1Z"/>`),
		g.Group(children),
	)
}