package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignRightRotateTwoHundredSeventyTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 19a1 1 0 0 1-1-1V3a1 1 0 0 1 2 0v15a1 1 0 0 1-1 1Zm12-4a1 1 0 0 1-1-1V3a1 1 0 1 1 2 0v11a1 1 0 0 1-1 1Zm-7 6a1 1 0 1 0 2 0V3a1 1 0 1 0-2 0v18Z"/>`),
		g.Group(children),
	)
}