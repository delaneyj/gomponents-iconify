package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotateOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" d="M44 24C44 12.9543 35.0457 4 24 4C12.9543 4 4 12.9543 4 24C4 35.0457 12.9543 44 24 44C30.9566 44 37.0836 40.4483 40.6667 35.0593"/><path stroke-linecap="round" d="M44 24H30"/><circle cx="24" cy="24" r="6" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}