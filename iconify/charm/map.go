package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Map(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.25 5.25v8.5m-4.5-10.5v8.5m-4 2.5v-9.5l4-2l4.5 2l4-2v9.5l-4 2l-4.5-2z"/>`),
		g.Group(children),
	)
}