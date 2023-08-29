package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiLaughSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 2a6 6 0 1 1 0 12A6 6 0 0 1 8 2ZM6 6a.5.5 0 0 1 .5.5a.5.5 0 0 0 1 0a1.5 1.5 0 1 0-3 0a.5.5 0 0 0 1 0A.5.5 0 0 1 6 6Zm4 0a.5.5 0 0 1 .5.5a.5.5 0 0 0 1 0a1.5 1.5 0 0 0-3 0a.5.5 0 0 0 1 0A.5.5 0 0 1 10 6ZM4.535 8.5a3.5 3.5 0 0 0 6.93 0h-6.93Z"/>`),
		g.Group(children),
	)
}