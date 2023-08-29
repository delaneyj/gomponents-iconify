package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBasket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 18a1 1 0 0 0 1-1v-2a1 1 0 0 0-2 0v2a1 1 0 0 0 1 1Zm-4 0a1 1 0 0 0 1-1v-2a1 1 0 0 0-2 0v2a1 1 0 0 0 1 1Zm9-12h-1.38l-1.73-3.45a1 1 0 1 0-1.78.9L15.38 6H8.62l1.27-2.55a1 1 0 0 0-1.78-.9L6.38 6H5a3 3 0 0 0-.92 5.84l.74 7.46a3 3 0 0 0 3 2.7h8.38a3 3 0 0 0 3-2.7l.74-7.46A3 3 0 0 0 19 6Zm-1.81 13.1a1 1 0 0 1-1 .9H7.81a1 1 0 0 1-1-.9L6.1 12h11.8ZM19 10H5a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2Z"/>`),
		g.Group(children),
	)
}