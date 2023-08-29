package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiSadSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 2a6 6 0 1 1 0 12A6 6 0 0 1 8 2ZM6.25 7.5a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm3.5 0a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm.118 3.322a.5.5 0 1 0 .764-.644c-1.325-1.57-3.94-1.57-5.264 0a.5.5 0 1 0 .764.644c.925-1.096 2.81-1.096 3.736 0Z"/>`),
		g.Group(children),
	)
}