package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8 14l-6 6h9v-3"/><path d="M18.37 3.63L8 14l3 3L21.37 6.63a2.12 2.12 0 1 0-3-3Z"/></g>`),
		g.Group(children),
	)
}