package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Agreement(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><rect width="32" height="40" x="8" y="4" stroke-linejoin="round" rx="2"/><path fill="#2F88FF" stroke-linejoin="round" d="M16 4H25V20L20.5 16L16 20V4Z"/><path d="M16 28H26"/><path d="M16 34H32"/></g>`),
		g.Group(children),
	)
}