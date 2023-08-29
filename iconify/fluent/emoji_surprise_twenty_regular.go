package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiSurpriseTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.5 9.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm4.25 3a1.75 1.75 0 1 1-3.5 0a1.75 1.75 0 0 1 3.5 0Zm.75-3a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM10 2a8 8 0 1 0 0 16a8 8 0 0 0 0-16Zm-7 8a7 7 0 1 1 14 0a7 7 0 0 1-14 0Z"/>`),
		g.Group(children),
	)
}