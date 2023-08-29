package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" fill-rule="evenodd" stroke="#000" d="M24 8L44 21V44H4L4 21L24 8Z" clip-rule="evenodd"/><path stroke="#fff" d="M20 44V23L12 28L12 44"/><path stroke="#fff" d="M28 44V23L36 28L36 44"/><path stroke="#000" d="M41 44H8"/></g>`),
		g.Group(children),
	)
}