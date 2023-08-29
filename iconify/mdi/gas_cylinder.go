package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GasCylinder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 9v11a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2V9c0-1.86 1.27-3.43 3-3.87V4H9V2h6v2h-2v1.13c1.73.44 3 2.01 3 3.87Z"/>`),
		g.Group(children),
	)
}