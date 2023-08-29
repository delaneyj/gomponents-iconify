package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FreezeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linejoin="round" d="M40 6H8C6.89543 6 6 6.89543 6 8V40C6 41.1046 6.89543 42 8 42H40C41.1046 42 42 41.1046 42 40V8C42 6.89543 41.1046 6 40 6Z"/><path stroke-linecap="round" d="M6 19.0586H42"/><path stroke-linecap="round" d="M16.1231 6L6 15"/><path stroke-linecap="round" d="M42 10.0068L32 19"/><path stroke-linecap="round" d="M26.123 6L11.9233 18.6242"/><path stroke-linecap="round" d="M36.123 6L21.9233 18.6242"/></g>`),
		g.Group(children),
	)
}