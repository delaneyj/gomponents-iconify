package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Qiyehao(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m28 19l-13.137-6.914C12.199 10.684 9 12.616 9 15.626V30"/><path d="m9 30l9-7v-9"/><path d="m11 12l9-7l16 8l-8 6m-8 10l13.137 6.914c2.664 1.402 5.863-.53 5.863-3.54V18"/><path d="m39 18l-9 7v9"/><path d="m37 36l-9 7l-16-8l8-6"/></g>`),
		g.Group(children),
	)
}