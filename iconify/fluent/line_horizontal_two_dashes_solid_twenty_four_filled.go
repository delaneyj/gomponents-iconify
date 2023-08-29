package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineHorizontalTwoDashesSolidTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 7a1 1 0 0 0 0 2h3a1 1 0 0 0 0-2H3Zm7.5 0a1 1 0 0 0 0 2h3a1 1 0 1 0 0-2h-3ZM18 7a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2h-3ZM3 15a1 1 0 1 0 0 2h18a1 1 0 1 0 0-2H3Z"/>`),
		g.Group(children),
	)
}