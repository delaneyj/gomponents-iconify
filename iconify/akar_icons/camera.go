package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m6.233 5.834l.445-2.226A2 2 0 0 1 8.64 2h6.72a2 2 0 0 1 1.962 1.608l.445 2.226a1.879 1.879 0 0 0 1.387 1.454A3.758 3.758 0 0 1 22 10.934V18a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4v-7.066a3.758 3.758 0 0 1 2.846-3.646a1.879 1.879 0 0 0 1.387-1.454Z"/><circle cx="12" cy="14" r="4"/><path d="M11 6h2"/></g>`),
		g.Group(children),
	)
}