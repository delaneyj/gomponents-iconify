package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M44 44V20L24 4L4 20L4 44H16V26H32V44H44Z"/><path stroke-linecap="round" d="M24 44V34"/></g>`),
		g.Group(children),
	)
}