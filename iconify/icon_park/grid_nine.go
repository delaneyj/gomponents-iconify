package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><rect width="38" height="38" x="5" y="5" stroke-linejoin="round" rx="2"/><path d="M5 18H43"/><path d="M5 30H43"/><path d="M17 5V43"/><path d="M30 5V43"/></g>`),
		g.Group(children),
	)
}