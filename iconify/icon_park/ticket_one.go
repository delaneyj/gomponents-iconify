package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TicketOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M4 8H11C11 8 12 13 17 13C22 13 23 8 23 8H44V40H23C23 40 22 35 17 35C12 35 11 40 11 40H4V8Z"/><path stroke="#fff" d="M17 19V21"/><path stroke="#fff" d="M17 27L17 29"/><path stroke="#fff" d="M25 21H36"/><path stroke="#fff" d="M25 27H36"/></g>`),
		g.Group(children),
	)
}