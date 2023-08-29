package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRotateCounterclockwiseTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 10a7 7 0 1 0-10 6.326V14.5a.5.5 0 0 1 1 0v3a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1 0-1h1.624A8 8 0 1 1 18 10a.5.5 0 0 1-1 0Zm-7 2a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0-1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}