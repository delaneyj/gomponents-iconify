package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallEndTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 2.999c-1.292 0-2.568.26-3.503.742c-.911.47-1.67 1.261-1.462 2.34l.101.65a1.5 1.5 0 0 0 1.632 1.262l.774-.077a1.5 1.5 0 0 0 1.332-1.262l.083-.535A4.117 4.117 0 0 1 6 6c.489 0 .836.064 1.044.119l.083.535a1.5 1.5 0 0 0 1.332 1.262l.773.077a1.5 1.5 0 0 0 1.633-1.261l.1-.652c.21-1.078-.55-1.87-1.46-2.34C8.567 3.259 7.291 3 6 3Z"/>`),
		g.Group(children),
	)
}