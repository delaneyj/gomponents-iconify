package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feBrush0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBrush1" fill="currentColor"><path id="feBrush2" d="M5.261 16.011A2 2 0 0 0 7.99 18.74A2.5 2.5 0 0 1 5.5 21H3v-2.5a2.5 2.5 0 0 1 2.261-2.489ZM19 3c1.1 0 2 1.006 2 2L8 18l-2-2L19 3Z"/></g></g>`),
		g.Group(children),
	)
}