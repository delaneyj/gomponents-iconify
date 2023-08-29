package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirstAidKit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M8 20v18a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V20M5 14h38v6H5v-6Zm26-6H17v6h14V8Z"/><path stroke-linecap="round" d="M20 30h8m-4-4v8"/></g>`),
		g.Group(children),
	)
}