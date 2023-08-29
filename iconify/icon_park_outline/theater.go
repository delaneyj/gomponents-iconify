package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Theater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="M8 6h32a2 2 0 0 1 2 2v28L32 22.005c-2.667 1.585-5.333 2.378-8 2.378s-5.333-.793-8-2.378L6 36V8a2 2 0 0 1 2-2Z" clip-rule="evenodd"/><path stroke-linecap="round" d="M6 42h36"/></g>`),
		g.Group(children),
	)
}