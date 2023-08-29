package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraphCuredIncreasing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.755 22.5h-21v-21"/><path d="M15.755 6a1 1 0 0 0-1-1h-2.5V2.5a1 1 0 0 0-1-1h-1.5a1 1 0 0 0-1 1V5h-2.5a1 1 0 0 0-1 1v1.5a1 1 0 0 0 1 1h2.5V11a1 1 0 0 0 1 1h1.5a1 1 0 0 0 1-1V8.5h2.5a1 1 0 0 0 1-1V6Zm-15 13.512h2.7c8.9 0 16.065-3.387 17.955-12.762"/><path d="m18.571 8.582l2.842-1.832l1.832 2.842"/></g>`),
		g.Group(children),
	)
}