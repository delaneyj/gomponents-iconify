package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedsideTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 18h40v24H4V18Zm18 6h4M4 30h40m-22 6h4M8 42v4m32-4v4M24 18v-8"/><path d="M32 10a8 8 0 1 0-16 0h16Z" clip-rule="evenodd"/><path d="M44 26v8M4 26v8"/></g>`),
		g.Group(children),
	)
}