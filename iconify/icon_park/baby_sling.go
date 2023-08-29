package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabySling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><circle cx="24" cy="10" r="5" fill="#2F88FF"/><path fill="#2F88FF" stroke-linecap="round" stroke-linejoin="round" d="M24 21C14 21 11 6 11 6L6 8L9 24.5C10.1667 25.1667 14.5 27.5 17 31C19.5 34.5 19 38 24.5 38C30 38 30.5 33.5 32 31C33.5 28.5 37.6667 25 39 24.5L42 8L37 6C37 6 34 21 24 21Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M34 29C34 29 41 33 41 42H36C36 36 31 33 31 33"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 29C15 29 8 33 8 42H13C13 36 18 33 18 33"/></g>`),
		g.Group(children),
	)
}