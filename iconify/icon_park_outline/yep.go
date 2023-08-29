package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yep(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M32 42c-4.736 2.706-11 2.5-16 0s-9-8-9-12.5c0-2.623 4.356-4.53 6.196-5.219c.46-.172.75-.64.677-1.126L11.586 7.909a3.404 3.404 0 1 1 6.75-.88L20 22l.83-15.77a3.408 3.408 0 1 1 6.795.518l-1.038 10.386A2.372 2.372 0 0 1 28.947 15A4.053 4.053 0 0 1 33 19.053V21.5a3.5 3.5 0 1 1 7 0v10.649c0 .566-.053 1.132-.253 1.661c-.596 1.577-2.46 5.169-7.747 8.19Z"/>`),
		g.Group(children),
	)
}