package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FreezeColumn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linejoin="round" d="M40 6H8C6.89543 6 6 6.89543 6 8V40C6 41.1046 6.89543 42 8 42H40C41.1046 42 42 41.1046 42 40V8C42 6.89543 41.1046 6 40 6Z"/><path stroke-linecap="round" d="M16.1231 6L6 15"/><path stroke-linecap="round" d="M19 12L6 24"/><path stroke-linecap="round" d="M19 21L6 33"/><path stroke-linecap="round" d="M19 30L6 42"/><path stroke-linecap="round" d="M19 6V42"/></g>`),
		g.Group(children),
	)
}