package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Expressvpn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.088 7.72H23.274a6.725 6.725 0 0 0-6.152 4.008l-2.19 4.96h23.269a.897.897 0 0 1 .82 1.259l-.693 1.569H8.46L4.5 28.484h4.894l-1.249 2.828l-.486 1.1a5.604 5.604 0 0 0 5.128 7.868h16.072a4.614 4.614 0 0 0 4.656-4.043a4.485 4.485 0 0 0-4.463-4.925H14.559a2.014 2.014 0 0 1-1.843-2.828h18.792a11.215 11.215 0 0 0 10.26-6.685l1.007-2.283A8.402 8.402 0 0 0 35.088 7.72Z"/>`),
		g.Group(children),
	)
}