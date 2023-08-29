package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareArrowForwardSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.5 4.5a2 2 0 0 1 2-2h7a2 2 0 0 1 2 2v1.39a5.5 5.5 0 0 0-7.61 7.61H4.5a2 2 0 0 1-2-2v-7Zm8 10.5a4.5 4.5 0 1 1 0-9a4.5 4.5 0 0 1 0 9Zm.896-6.396l.897.896H10.75A2.75 2.75 0 0 0 8 12.25v.25a.5.5 0 0 0 1 0v-.25c0-.966.784-1.75 1.75-1.75h1.543l-.897.896a.5.5 0 0 0 .708.708l1.752-1.753a.499.499 0 0 0-.002-.705l-1.75-1.75a.5.5 0 0 0-.708.708Z"/>`),
		g.Group(children),
	)
}