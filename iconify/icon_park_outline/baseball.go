package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Baseball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path d="M8.546 11.273C12.788 14.909 14.91 19.15 14.91 24s-2.122 9.09-6.364 12.727m30.909 0C35.212 33.091 33.09 28.85 33.09 24s2.121-9.09 6.364-12.727"/></g>`),
		g.Group(children),
	)
}