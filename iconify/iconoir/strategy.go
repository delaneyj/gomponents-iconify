package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Strategy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6 20.5C7 11 11.5 8 20 6"/><path d="m15.909 3.81l4.486 2.09l-2.092 4.486M5 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm11 13.243l2.121-2.122m0 0L20.243 16m-2.122 2.121L16 16m2.121 2.121l2.122 2.122"/></g>`),
		g.Group(children),
	)
}