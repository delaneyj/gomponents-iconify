package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraphCuredDecreasing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.777 22.5h-21v-21m21 4.5a1 1 0 0 0-1-1h-2.5V2.5a1 1 0 0 0-1-1h-1.5a1 1 0 0 0-1 1V5h-2.5a1 1 0 0 0-1 1v1.5a1 1 0 0 0 1 1h2.5V11a1 1 0 0 0 1 1h1.5a1 1 0 0 0 1-1V8.5h2.5a1 1 0 0 0 1-1V6Z"/><path d="M.777 9.5h1.4a10.826 10.826 0 0 1 8.6 4.25a10.824 10.824 0 0 0 8.6 4.25h3.846"/><path d="M20.833 15.609L23.223 18l-2.39 2.391"/></g>`),
		g.Group(children),
	)
}