package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignRightRotateNinetyTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 5a1 1 0 0 1 1 1v15a1 1 0 1 1-2 0V6a1 1 0 0 1 1-1ZM6 9a1 1 0 0 1 1 1v11a1 1 0 1 1-2 0V10a1 1 0 0 1 1-1Zm7-6a1 1 0 1 0-2 0v18a1 1 0 1 0 2 0V3Z"/>`),
		g.Group(children),
	)
}