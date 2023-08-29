package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RainyFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.86 18l-3.153-3.153a1 1 0 0 0-1.414 0L8.18 17.96a8.001 8.001 0 1 1 7.8-11.873A6 6 0 1 1 17 18h-1.139Zm-5.628.732L12 16.965l1.768 1.767a2.5 2.5 0 1 1-3.536 0Z"/>`),
		g.Group(children),
	)
}