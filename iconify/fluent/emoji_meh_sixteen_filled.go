package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiMehSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 8a6 6 0 1 0 12 0A6 6 0 0 0 2 8Zm5-1a.75.75 0 1 1-1.5 0A.75.75 0 0 1 7 7Zm3.5 0A.75.75 0 1 1 9 7a.75.75 0 0 1 1.5 0ZM6 9h4a.5.5 0 0 1 0 1H6a.5.5 0 0 1 0-1Z"/>`),
		g.Group(children),
	)
}