package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitCherryPick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="5" cy="8" r="2.25"/><path d="M5 10.75v3.5m0-12.5v3.5M11.75 8h1.5m-4.5-3.25h1.5l1 3.25l-1 3.25h-1.5"/></g>`),
		g.Group(children),
	)
}