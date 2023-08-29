package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlidingHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M24 44C35.0457 44 44 35.0457 44 24C44 12.9543 35.0457 4 24 4C12.9543 4 4 12.9543 4 24C4 35.0457 12.9543 44 24 44Z"/><path stroke="#fff" stroke-linecap="round" d="M28 20L32 24L28 28"/><path stroke="#fff" stroke-linecap="round" d="M20 20L16 24L20 28"/></g>`),
		g.Group(children),
	)
}