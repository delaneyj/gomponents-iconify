package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareCircleLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M12 9a3 3 0 1 1 0-6a3 3 0 0 1 0 6ZM5.5 21a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm13 0a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/><path stroke-linecap="round" d="M20 13a7.98 7.98 0 0 0-2.708-6M4 13a7.98 7.98 0 0 1 2.708-6M10 20.748A8 8 0 0 0 12 21a8 8 0 0 0 2-.252"/></g>`),
		g.Group(children),
	)
}