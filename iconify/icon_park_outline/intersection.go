package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Intersection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M31 31h9a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2H19a2 2 0 0 0-2 2v9"/><path d="M17 17H8a2 2 0 0 0-2 2v21a2 2 0 0 0 2 2h21a2 2 0 0 0 2-2v-9"/><rect width="14" height="14" x="17" y="17" rx="2"/></g>`),
		g.Group(children),
	)
}