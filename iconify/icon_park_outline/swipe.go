package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Swipe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m6 15l13.371 10.538C20.5 26.5 22.828 28 25 26c2.29-2.108.5-4.5 0-5l-8-6M4 8h23l11 10M9 33l35 .02"/><path d="M9 18v22h35V18H22"/></g>`),
		g.Group(children),
	)
}