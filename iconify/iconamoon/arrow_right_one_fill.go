package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightOneFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 19a1 1 0 0 0 1.707.707l7-7a1 1 0 0 0 0-1.414l-7-7A1 1 0 0 0 11 5v6H5a1 1 0 1 0 0 2h6v6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}