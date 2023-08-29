package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TicketsOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="26" height="38" x="5" y="42" fill="#2F88FF" stroke="#000" stroke-linejoin="bevel" stroke-width="4" rx="2" transform="rotate(-90 5 42)"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M9.00002 16L32 4.99999L37 16"/><circle cx="13" cy="23" r="2" fill="#fff"/><circle cx="13" cy="29" r="2" fill="#fff"/><circle cx="13" cy="35" r="2" fill="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M21 35H25L36 23"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 29H30"/></g>`),
		g.Group(children),
	)
}