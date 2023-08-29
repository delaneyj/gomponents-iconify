package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCardRemoveFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 3a1 1 0 0 1 1-1h2a1 1 0 0 1 .997.923L6.157 5H21a1 1 0 0 1 .994 1.11l-.816 7.345a3 3 0 0 1-2.768 2.661l-10.355.74a3 3 0 0 1-3.205-2.763l-.616-8.016L4.074 4H3a1 1 0 0 1-1-1Zm9.5 7a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2h-3ZM16 20.5a1.5 1.5 0 0 1 1.5-1.5h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5h-.01a1.5 1.5 0 0 1-1.5-1.5v-.01ZM8.5 19A1.5 1.5 0 0 0 7 20.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5v-.01a1.5 1.5 0 0 0-1.5-1.5H8.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}