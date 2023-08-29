package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HangerTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M37 32h3.91a3.09 3.09 0 0 0 1.382-5.854L24 17L5.708 26.146A3.09 3.09 0 0 0 7.09 32H11"/><path d="M11 30h26v14H11V30Zm13-13s4-6.79 4-9a4 4 0 0 0-8 0"/></g>`),
		g.Group(children),
	)
}