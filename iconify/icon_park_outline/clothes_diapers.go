package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesDiapers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 11h36v8s0 8-4 12s-10.158 6-10.158 6h-7.684S14 35 10 31c-4-4-4-12-4-12v-8Z"/><path d="M20.158 37s.1-7.075-3.158-11c-3.044-3.669-11-7-11-7m21.842 18s-.1-7.075 3.158-11c3.044-3.669 11-7 11-7"/></g>`),
		g.Group(children),
	)
}