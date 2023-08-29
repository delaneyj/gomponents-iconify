package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ticket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><path stroke="#000" stroke-linejoin="round" d="M9.00013 16.0001L34 6.00008L38.0004 16.0001"/><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M4 16H44V22C41 22 38 24 38 27.5C38 31 41 34 44 34V40H4V34C7.00016 34 10 32 10 28C10 24 7 22 4 22V16Z"/><path stroke="#fff" d="M17 25.3848H23"/><path stroke="#fff" d="M17 31.3848H31"/></g>`),
		g.Group(children),
	)
}