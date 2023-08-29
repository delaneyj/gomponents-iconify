package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M6 44V25h6v-8h24v8h6v19H6Z"/><path stroke-linecap="round" d="M17 17V8l14-4v13"/></g>`),
		g.Group(children),
	)
}