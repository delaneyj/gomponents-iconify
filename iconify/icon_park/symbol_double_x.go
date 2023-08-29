package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SymbolDoubleX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M10 4L24 18L38 4"/><path d="M10 44L24 30L38 44"/><path d="M43 10L29 24L43 38"/><path d="M4 10L18 24L4 38"/></g>`),
		g.Group(children),
	)
}