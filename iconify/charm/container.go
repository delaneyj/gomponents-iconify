package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Container(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m1.75 12.2l5.5 2l7-4.5v-6l-5.5-2l-7 4.5z"/><path d="M10.8 6.25v5.5m-3.5-3.5v6m-5.5-8l5.5 2l7-4.5"/></g>`),
		g.Group(children),
	)
}