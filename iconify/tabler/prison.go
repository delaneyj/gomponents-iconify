package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Prison(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 4v16M14 4v16M6 4v5m0 6v5m4-16v5m1 0H5v6h6zm-1 6v5m-2-8h-.01"/>`),
		g.Group(children),
	)
}