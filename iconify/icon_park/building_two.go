package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M21 13L11 20L11 44"/><path fill="#2F88FF" fill-rule="evenodd" d="M21 4L31 11V24L38 29V44H21V4Z" clip-rule="evenodd"/><path d="M4 44H44"/></g>`),
		g.Group(children),
	)
}