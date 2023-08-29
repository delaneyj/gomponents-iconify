package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCartThreeShoppingCartCheckout(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m3.75 4.5l.25 1m3.25-1L7 5.5m-5.18 2H10l1-5H1a.5.5 0 0 0-.49.59l.82 4a.49.49 0 0 0 .49.41Zm9.18-5l.42-1.6a.5.5 0 0 1 .49-.4h1.59m-3.5 7l-.42 2.1a.5.5 0 0 1-.49.4H3"/><circle cx="3.5" cy="13" r=".5"/><circle cx="8.5" cy="13" r=".5"/></g>`),
		g.Group(children),
	)
}