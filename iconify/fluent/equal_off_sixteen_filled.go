package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualOffSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m10.293 11l3.853 3.854a.5.5 0 0 0 .708-.708l-13-13a.5.5 0 1 0-.708.708L3.793 4.5H2.75a.75.75 0 0 0 0 1.5h2.543l3.5 3.5H2.75a.75.75 0 0 0 0 1.5h7.543Zm1.328-1.5l1.5 1.5h.129a.75.75 0 0 0 0-1.5h-1.629Zm-5-5l1.5 1.5h5.129a.75.75 0 0 0 0-1.5H6.621Z"/>`),
		g.Group(children),
	)
}