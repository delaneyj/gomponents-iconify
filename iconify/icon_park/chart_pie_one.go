package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPieOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M5 24C5 34.4934 13.5066 43 24 43V26C24 24.8954 24.8954 24 26 24H43C43 13.5066 34.4934 5 24 5C13.5066 5 5 13.5066 5 24Z"/><rect width="12" height="12" x="30" y="30"/></g>`),
		g.Group(children),
	)
}