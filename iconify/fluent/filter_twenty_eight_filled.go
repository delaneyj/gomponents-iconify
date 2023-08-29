package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 19a1 1 0 1 1 0 2h-6a1 1 0 1 1 0-2h6Zm4-6a1 1 0 1 1 0 2H7a1 1 0 1 1 0-2h14Zm3-6a1 1 0 1 1 0 2H4a1 1 0 0 1 0-2h20Z"/>`),
		g.Group(children),
	)
}