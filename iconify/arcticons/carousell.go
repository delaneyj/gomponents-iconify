package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Carousell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M28.88 37.614a13.116 13.116 0 1 1-.104-24.39"/><path d="M27.112 32.66a7.862 7.862 0 1 1-.065-14.467m.207 14.493l1.621 4.898m-.062-24.307l-1.559 4.903m2.249 1.051l2.984-3.258m-.844 6.175l3.964-1.948m-6.104 11.506l2.984 3.258m-.844-6.174l3.964 1.948m-3.357-5.253h4.415"/><path d="M37.73 4.732c0-.128-.161-.232-.363-.232h-6.013c-.201 0-.366.104-.366.232V7.49H7.94c-1.075 0-1.947.87-1.947 1.945v32.118c0 1.075.872 1.947 1.947 1.947h32.12a1.947 1.947 0 0 0 1.947-1.947V9.433a1.947 1.947 0 0 0-1.947-1.946l-2.33.003V4.732Z"/></g>`),
		g.Group(children),
	)
}