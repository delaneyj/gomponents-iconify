package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RectangleTear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M6 8V40C6 41.1046 6.89543 42 8 42H40C41.1046 42 42 41.1046 42 40V8C42 6.89543 41.1046 6 40 6H8C6.89543 6 6 6.89543 6 8Z"/><path stroke="#fff" d="M27 6L21 12L27 18L21 24L27 30L21 36L27 42"/><path stroke="#000" d="M18 6H34"/><path stroke="#000" d="M18 42H34"/></g>`),
		g.Group(children),
	)
}