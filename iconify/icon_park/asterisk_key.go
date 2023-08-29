package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AsteriskKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" rx="3"/><path stroke="#fff" d="M24 16V32"/><path stroke="#fff" d="M17.447 19.4114L30.5535 28.5886"/><path stroke="#fff" d="M30.5532 19.4114L17.4468 28.5886"/></g>`),
		g.Group(children),
	)
}