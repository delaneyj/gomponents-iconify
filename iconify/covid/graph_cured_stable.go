package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraphCuredStable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.787 22.5h-21v-21"/><path d="M16.537 6a1 1 0 0 0-1-1h-2.5V2.5a1 1 0 0 0-1-1h-1.5a1 1 0 0 0-1 1V5h-2.5a1 1 0 0 0-1 1v1.5a1 1 0 0 0 1 1h2.5V11a1 1 0 0 0 1 1h1.5a1 1 0 0 0 1-1V8.5h2.5a1 1 0 0 0 1-1V6Zm4.285 9.234l2.391 2.391l-2.391 2.391"/><path d="M.787 17.625h.5a10.6 10.6 0 0 0 4.737-1.118L8.446 15.3a6.352 6.352 0 0 1 5.683 0l2.422 1.211a10.6 10.6 0 0 0 4.736 1.118h1.926"/></g>`),
		g.Group(children),
	)
}