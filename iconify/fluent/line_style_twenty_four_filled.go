package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineStyleTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 5a1 1 0 0 0 0 2h3a1 1 0 0 0 0-2H3Zm7.5 0a1 1 0 0 0 0 2h3a1 1 0 1 0 0-2h-3ZM18 5a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2h-3ZM2 12a1 1 0 0 1 1-1h18a1 1 0 1 1 0 2H3a1 1 0 0 1-1-1Zm1.5 5a1.5 1.5 0 0 0 0 3h17a1.5 1.5 0 0 0 0-3h-17Z"/>`),
		g.Group(children),
	)
}