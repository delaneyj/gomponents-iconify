package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForFlagBurkinaFaso(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#009E49" d="M19.602 18.942l.99 3.048L18 20.106l-2.593 1.884l.99-3.048L15.101 18H0v9a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4v-9H20.899l-1.297.942z"/><path fill="#EF2B2D" d="M32 5H4a4 4 0 0 0-4 4v9h15.101l-1.296-.942h3.205L18 14.01l.99 3.048h3.205L20.899 18H36V9a4 4 0 0 0-4-4z"/><path fill="#FCD116" d="M15.407 21.99L18 20.106l2.593 1.884l-.991-3.048L20.899 18l1.296-.942H18.99L18 14.01l-.99 3.048h-3.205l1.296.942l1.297.942z"/>`),
		g.Group(children),
	)
}