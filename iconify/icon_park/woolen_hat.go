package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WoolenHat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><rect width="40" height="10" x="4" y="34" fill="#2F88FF" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 26V34"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 26V34"/><path stroke-linecap="round" stroke-linejoin="round" d="M33 26V34"/><circle cx="24" cy="8" r="4" fill="#2F88FF"/><path d="M8 34C8 25.75 9 12 24 12C39 12 40 25.75 40 34"/></g>`),
		g.Group(children),
	)
}