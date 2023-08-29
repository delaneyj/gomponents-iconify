package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M44 24C44 12.9543 35.0457 4 24 4C12.9543 4 4 12.9543 4 24H44Z"/><path stroke-linecap="round" d="M24 24V38.5536C24 41.5678 26.4858 44.0112 29.5 44.0112C32.5142 44.0112 35 41.5678 35 38.5536"/></g>`),
		g.Group(children),
	)
}