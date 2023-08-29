package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiMehTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 10a8 8 0 1 1 16 0a8 8 0 0 1-16 0Zm6.5-1.5a1 1 0 1 0-2 0a1 1 0 0 0 2 0Zm4 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-6.5 3a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 0-1h-7a.5.5 0 0 0-.5.5Z"/>`),
		g.Group(children),
	)
}