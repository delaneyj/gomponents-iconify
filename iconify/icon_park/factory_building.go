package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FactoryBuilding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M4 44V4H12V20L28 12V20L44 12V44H4Z"/><rect width="8" height="8" x="12" y="28" fill="#43CCF8" stroke="#fff"/><rect width="8" height="8" x="28" y="28" fill="#43CCF8" stroke="#fff"/></g>`),
		g.Group(children),
	)
}