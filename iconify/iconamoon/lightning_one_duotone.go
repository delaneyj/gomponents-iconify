package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightningOneDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="m6 14l7-12v8h5l-7 12v-8H6Z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="m6 14l7-12v8h5l-7 12v-8H6Z"/></g>`),
		g.Group(children),
	)
}