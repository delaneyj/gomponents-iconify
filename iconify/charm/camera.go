package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 4.75v8.5h12.5v-8.5h-3l-1.5-2h-3.5l-1.5 2z"/><circle cx="8" cy="8.5" r="2.25"/></g>`),
		g.Group(children),
	)
}