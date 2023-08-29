package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualOffSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m10.293 11l3.853 3.854a.5.5 0 0 0 .708-.708l-13-13a.5.5 0 1 0-.708.708L4.293 5H2.5a.5.5 0 0 0 0 1h2.793l4 4H2.5a.5.5 0 0 0 0 1h7.793Zm1.828-1l1 1h.379a.5.5 0 0 0 0-1h-1.379Zm-5-5l1 1H13.5a.5.5 0 0 0 0-1H7.121Z"/>`),
		g.Group(children),
	)
}