package heroicons_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReceiptTax(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 14l6-6m-5.5.5h.01m4.99 5h.01M19 21V5a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v16l3.5-2l3.5 2l3.5-2l3.5 2ZM10 8.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm5 5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z"/>`),
		g.Group(children),
	)
}