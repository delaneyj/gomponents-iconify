package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DroneOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M29 18V16C29 13.2386 26.7614 11 24 11V11C21.2386 11 19 13.2386 19 16V18"/><path fill="#2F88FF" stroke-linecap="round" d="M17 18H31L28.8462 25H19.1538L17 18Z"/><rect width="7" height="7" x="5" y="22" fill="#2F88FF"/><rect width="7" height="7" x="36" y="22" fill="#2F88FF"/><path stroke-linecap="round" d="M16 8L4 8"/><path stroke-linecap="round" d="M30 33L34 40"/><path stroke-linecap="round" d="M18 33L14 40"/><path stroke-linecap="round" d="M44 8L32 8"/></g>`),
		g.Group(children),
	)
}