package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualOffTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.854 2.146a.5.5 0 1 0-.708.708L5.793 6.5H3.75a.75.75 0 0 0 0 1.5h3.543l3.5 3.5H3.75a.75.75 0 0 0 0 1.5h8.543l4.853 4.854a.5.5 0 0 0 .708-.708l-15-15ZM13.62 11.5l1.5 1.5h1.129a.75.75 0 0 0 0-1.5H13.62Zm-5-5l1.5 1.5h6.129a.75.75 0 0 0 0-1.5H8.621Z"/>`),
		g.Group(children),
	)
}