package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Joystick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M44 33H4V43H44V33Z"/><path stroke-linecap="round" d="M38 26H26V33H38V26Z"/><path fill="#2F88FF" d="M18 14C20.7614 14 23 11.7614 23 9C23 6.23858 20.7614 4 18 4C15.2386 4 13 6.23858 13 9C13 11.7614 15.2386 14 18 14Z"/><path stroke-linecap="round" d="M16 14L9 33"/></g>`),
		g.Group(children),
	)
}