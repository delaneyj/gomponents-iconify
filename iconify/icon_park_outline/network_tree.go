package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkTree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 34h8v8H4zM8 6h32v12H8zm16 28V18"/><path d="M8 34v-8h32v8m-4 0h8v8h-8zm-16 0h8v8h-8zm-6-22h2"/></g>`),
		g.Group(children),
	)
}