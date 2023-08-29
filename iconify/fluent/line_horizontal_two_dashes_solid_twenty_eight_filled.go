package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineHorizontalTwoDashesSolidTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 9a1 1 0 0 0 0 2h4a1 1 0 1 0 0-2H3Zm9 0a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2h-4Zm9 0a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2h-4ZM3 17a1 1 0 1 0 0 2h22a1 1 0 1 0 0-2H3Z"/>`),
		g.Group(children),
	)
}