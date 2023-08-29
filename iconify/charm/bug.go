package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="10" r="4.25"/><path d="M14.25 10.25h-1.5m-1 2.5l1.5 1.5m0-8.5l-1.5 1.5m-10 3h1.5m1 2.5l-1.5 1.5m0-8.5l1.5 1.5m1.5-1.5s-.75-3 2.25-3s2.25 3 2.25 3"/></g>`),
		g.Group(children),
	)
}