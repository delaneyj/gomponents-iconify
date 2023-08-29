package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Code(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M16 13L4 25.4322L16 37"/><path stroke-linejoin="round" d="M32 13L44 25.4322L32 37"/><path d="M28 4L21 44"/></g>`),
		g.Group(children),
	)
}