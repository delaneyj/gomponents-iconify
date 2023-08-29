package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentOffTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.28 2.22a.75.75 0 1 0-1.06 1.06l.702.702A3.24 3.24 0 0 0 2 6.25v8.5A3.25 3.25 0 0 0 5.25 18H6v2.75c0 1.03 1.176 1.618 2 1L13 18h3.939l3.78 3.78a.75.75 0 0 0 1.061-1.06L3.28 2.22ZM22 14.75a3.246 3.246 0 0 1-1.398 2.67L6.182 3H18.75A3.25 3.25 0 0 1 22 6.25v8.5Z"/>`),
		g.Group(children),
	)
}