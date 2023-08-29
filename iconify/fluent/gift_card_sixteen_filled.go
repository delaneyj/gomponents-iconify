package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GiftCardSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 3h2v2.085A1.5 1.5 0 0 0 3.085 7H1V5a2 2 0 0 1 2-2Zm3 5.707l1.146 1.147a.5.5 0 0 0 .708-.708L6.707 8H15v3a2 2 0 0 1-2 2H6V8.707ZM4.293 8L3.146 9.146a.5.5 0 1 0 .708.708L5 8.707V13H3a2 2 0 0 1-2-2V8h3.293Zm3.622-1A1.5 1.5 0 0 0 6 5.085V3h7a2 2 0 0 1 2 2v2H7.915ZM6 6.5V7h.5a.5.5 0 1 0-.5-.5Zm-1 0V7h-.5a.5.5 0 1 1 .5-.5Z"/>`),
		g.Group(children),
	)
}