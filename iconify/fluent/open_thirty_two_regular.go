package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.5 5A3.5 3.5 0 0 0 5 8.5v15A3.5 3.5 0 0 0 8.5 27h15a3.5 3.5 0 0 0 3.5-3.5V19a1 1 0 1 1 2 0v4.5a5.5 5.5 0 0 1-5.5 5.5h-15A5.5 5.5 0 0 1 3 23.5v-15A5.5 5.5 0 0 1 8.5 3H13a1 1 0 1 1 0 2H8.5ZM18 4a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1v9a1 1 0 1 1-2 0V6.414l-7.293 7.293a1 1 0 0 1-1.414-1.414L25.586 5H19a1 1 0 0 1-1-1Z"/>`),
		g.Group(children),
	)
}