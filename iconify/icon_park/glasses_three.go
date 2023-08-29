package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlassesThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><circle cx="12" cy="33" r="7" fill="#2F88FF"/><path stroke-linecap="round" d="M29 33C29 29.6863 27.5 27 24 27C20.5 27 19 29.6863 19 33"/><circle cx="36" cy="33" r="7" fill="#2F88FF"/><path stroke-linecap="round" d="M14 8H12C8.68629 8 6 10.6863 6 14V22"/><path stroke-linecap="round" d="M34 8H36C39.3137 8 42 10.6863 42 14V22"/></g>`),
		g.Group(children),
	)
}