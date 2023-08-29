package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReceiptTextLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7h8m-8 4h8m-8 4h4m8 6V5a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v16l2.5-2l3 2l2.5-2l2.5 2l3-2l2.5 2z"/>`),
		g.Group(children),
	)
}