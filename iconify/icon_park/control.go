package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Control(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><rect width="30" height="40" x="9" y="4" rx="2"/><circle cx="24" cy="31" r="6" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M18 12H21"/><path stroke-linecap="round" stroke-linejoin="round" d="M18 18H21"/><path stroke-linecap="round" stroke-linejoin="round" d="M27 18H30"/></g>`),
		g.Group(children),
	)
}