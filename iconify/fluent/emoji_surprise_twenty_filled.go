package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiSurpriseTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 10a8 8 0 1 1 16 0a8 8 0 0 1-16 0Zm5.5-.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm4.25 3a1.75 1.75 0 1 0-3.5 0a1.75 1.75 0 0 0 3.5 0Zm.75-3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}