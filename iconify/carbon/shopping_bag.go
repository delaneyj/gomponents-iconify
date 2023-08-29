package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28.76 11.35A1 1 0 0 0 28 11h-6V7a3 3 0 0 0-3-3h-6a3 3 0 0 0-3 3v4H4a1 1 0 0 0-1 1.15L4.88 24.3a2 2 0 0 0 2 1.7h18.26a2 2 0 0 0 2-1.7L29 12.15a1 1 0 0 0-.24-.8ZM12 7a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v4h-8Zm13.14 17H6.86L5.17 13h21.66Z"/>`),
		g.Group(children),
	)
}