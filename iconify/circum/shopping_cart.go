package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.437 19.934a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM6.22 15.668l-.945-10.9a.75.75 0 0 0-.749-.693H3.56a.5.5 0 0 1 0-1h.966a1.75 1.75 0 0 1 1.746 1.617l.139 1.818h13.03c.885 0 1.577.76 1.494 1.638l-.668 7.52a2.5 2.5 0 0 1-2.489 2.267H8.709a2.5 2.5 0 0 1-2.489-2.267Zm.274-8.158l.722 8.066a1.5 1.5 0 0 0 1.493 1.359h9.069a1.5 1.5 0 0 0 1.493-1.359l.668-7.518a.5.5 0 0 0-.498-.548H6.494Zm4.454 12.424a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/>`),
		g.Group(children),
	)
}