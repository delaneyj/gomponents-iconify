package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BreastPump(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M35 25C35 21 29 19 29 19V14H17V19C17 19 11 21 11 25V44H35V25Z"/><path d="M20 4L13 10"/><path d="M23 14L17 7"/><path d="M26 8H35V15L41 20"/></g>`),
		g.Group(children),
	)
}