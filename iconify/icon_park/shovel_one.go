package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShovelOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M11 4H37L34.1667 24H13.8333L11 4Z"/><path stroke="#fff" d="M21 11L21 17"/><path stroke="#fff" d="M27 11L27 17"/><path stroke="#000" d="M24 24V44"/></g>`),
		g.Group(children),
	)
}