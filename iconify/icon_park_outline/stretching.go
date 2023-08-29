package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stretching(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M23 6H8a2 2 0 0 0-2 2v32a2 2 0 0 0 2 2h32a2 2 0 0 0 2-2V25"/><path stroke-linecap="round" d="M24.001 16v8M42 6v8m-9.999 10h-8"/><path d="M42 6L24 24"/><path stroke-linecap="round" d="M42 6h-8"/></g>`),
		g.Group(children),
	)
}