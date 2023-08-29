package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Card(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M23.5 7.25c-3-.5-8.5-.75-11.5-.75s-8.5.25-11.5.75m23 4c-3-.5-8.5-.75-11.5-.75s-8.5.25-11.5.75m12 5.75c0-1.652.607-2.5 1.618-2.5c1.012 0 1.382 1.273 1.382 2.121h.047s.286-.848 1.203-.848s.991.943.991.943c.897.05 1.784.11 2.759.185m3-12.651V20.5H23c-3-.5-8-.75-11-.75S4 20 1 20.5H.5V4.25c3-.5 8.5-.75 11.5-.75s8.5.25 11.5.75Z"/>`),
		g.Group(children),
	)
}