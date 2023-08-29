package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HanfuChineseStyle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M9 9L18 4H30L39 9L43 25L35 29V44H13V29L5 25L9 9Z"/><path d="M18 4L24 14.5"/><path d="M30 4L24 14.5L13 29"/></g>`),
		g.Group(children),
	)
}