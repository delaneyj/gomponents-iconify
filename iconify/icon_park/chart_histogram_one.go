package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartHistogramOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 6V42H42"/><path d="M18 34H14"/><path d="M26 26H14"/><path d="M42 18H14"/><path d="M34 10L14 10"/></g>`),
		g.Group(children),
	)
}