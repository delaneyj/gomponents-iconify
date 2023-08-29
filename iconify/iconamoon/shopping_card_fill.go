package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCardFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 2a1 1 0 0 0 0 2h1.074l.16 2.077l.616 8.017a3 3 0 0 0 3.205 2.762l10.355-.74a3 3 0 0 0 2.768-2.661l.816-7.345A1 1 0 0 0 21 5H6.157l-.16-2.077A1 1 0 0 0 5 2H3Zm4 18.5A1.5 1.5 0 0 1 8.5 19h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H8.5a1.5 1.5 0 0 1-1.5-1.5v-.01ZM17.5 19a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5v-.01a1.5 1.5 0 0 0-1.5-1.5h-.01Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}