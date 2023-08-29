package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tangerine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.786 8.346L7.398 8.344c-1.823 0-3.161 2.636.02 3.638c9.46 2.638 17.665 10.643 14.086 25.085c-.587 2.193 1.796 4.161 4.049.736c1.82-3.154 16.563-25.919 16.563-25.919c.898-1.564.261-3.539-2.329-3.539h0Z"/>`),
		g.Group(children),
	)
}