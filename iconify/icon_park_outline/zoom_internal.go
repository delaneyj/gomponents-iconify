package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomInternal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M44 4H4v40h40V4Z"/><path stroke-linecap="round" d="M16 4v12H4m32 8v12H24m12 0L24 24M4 6v20M7 4h20"/></g>`),
		g.Group(children),
	)
}