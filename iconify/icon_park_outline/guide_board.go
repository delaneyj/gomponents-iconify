package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GuideBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 4v37"/><path d="M24 8h15.545L42 12l-2.455 4H24V8Zm0 14H8.455L6 26l2.455 4H24v-8Z"/><path stroke-linecap="round" d="M16 42h16"/></g>`),
		g.Group(children),
	)
}