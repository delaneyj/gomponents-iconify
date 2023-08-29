package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignLeftRotateNinetyTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 2a1 1 0 0 1 1 1v15a1 1 0 1 1-2 0V3a1 1 0 0 1 1-1ZM6 2a1 1 0 0 1 1 1v11a1 1 0 1 1-2 0V3a1 1 0 0 1 1-1Zm7 1a1 1 0 1 0-2 0v18a1 1 0 1 0 2 0V3Z"/>`),
		g.Group(children),
	)
}