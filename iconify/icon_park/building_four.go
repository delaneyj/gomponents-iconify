package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" fill-rule="evenodd" stroke="#000" d="M17 14L44 24V44H17L17 14Z" clip-rule="evenodd"/><path stroke="#000" d="M17 14L4 24L4 44H17"/><path stroke="#fff" d="M35 44V32L26 29L26 44"/><path stroke="#000" d="M44 44H17"/></g>`),
		g.Group(children),
	)
}