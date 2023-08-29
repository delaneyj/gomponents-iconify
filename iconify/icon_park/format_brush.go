package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatBrush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M34 5H6V20H34V5Z"/><path stroke-linecap="round" d="M34.0251 12H43V28.1014L19 31.2004V43"/></g>`),
		g.Group(children),
	)
}