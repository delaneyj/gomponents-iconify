package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M14 11L4 24L14 37H44V11H14Z"/><path stroke="#fff" d="M21 19L31 29"/><path stroke="#fff" d="M31 19L21 29"/></g>`),
		g.Group(children),
	)
}