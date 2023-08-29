package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronsDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.071 5.314l4.95-4.95a1 1 0 1 1 1.414 1.414L7.778 7.435a1 1 0 0 1-1.414 0L.707 1.778A1 1 0 1 1 2.121.364l4.95 4.95zm0 6l4.95-4.95a1 1 0 1 1 1.414 1.414l-5.657 5.657a1 1 0 0 1-1.414 0L.707 7.778a1 1 0 1 1 1.414-1.414l4.95 4.95z"/>`),
		g.Group(children),
	)
}