package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterRateTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" fill-rule="evenodd" stroke="#000" d="M24 44C32.8366 44 40 36.8363 40 28C40 15 24 4 24 4C24 4 8 15 8 28C8 36.8366 15.1634 44 24 44Z" clip-rule="evenodd"/><path stroke="#fff" d="M24 20L20 28H28L24 36"/></g>`),
		g.Group(children),
	)
}