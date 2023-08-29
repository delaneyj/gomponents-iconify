package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignLeftRotateTwoHundredSeventyTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 22a1 1 0 0 1-1-1V6a1 1 0 0 1 2 0v15a1 1 0 0 1-1 1Zm12 0a1 1 0 0 1-1-1V10a1 1 0 1 1 2 0v11a1 1 0 0 1-1 1Zm-7-1a1 1 0 1 0 2 0V3a1 1 0 1 0-2 0v18Z"/>`),
		g.Group(children),
	)
}