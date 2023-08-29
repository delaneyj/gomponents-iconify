package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataSheet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="32" x="4" y="8" fill="#2F88FF" stroke="#000" rx="2"/><path stroke="#fff" d="M32 25V32"/><path stroke="#fff" d="M24 16V32"/><path stroke="#fff" d="M16 20V32"/></g>`),
		g.Group(children),
	)
}