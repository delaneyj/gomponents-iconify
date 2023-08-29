package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Baremetrics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22.109 7.951l1.485 2.464a3.507 3.507 0 0 1 0 3.275l-4.505 7.717a3.333 3.333 0 0 1-2.94 1.793H7.83a3.335 3.335 0 0 1-2.94-1.793l-1.555-2.632l6.139-5.695l4.447 2.578a1.093 1.093 0 0 0 1.456-.198zm-13.39.628L1.99 16.15L.406 13.725a3.495 3.495 0 0 1 0-3.27L5.158 2.59A3.338 3.338 0 0 1 8.1.8h8.008c1.228 0 2.357.687 2.942 1.79l1.616 2.722l-6.017 5.592l-4.432-2.574a1.098 1.098 0 0 0-1.499.248z"/>`),
		g.Group(children),
	)
}