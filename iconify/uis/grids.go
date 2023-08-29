package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grids(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 3.9v16.2c0 1 .9 1.9 1.9 1.9H8V2H3.9C2.9 2 2 2.9 2 3.9zM20.1 2H16v20h4.1c1 0 1.9-.9 1.9-1.9V3.9c0-1-.9-1.9-1.9-1.9zM10 22h4V2h-4v20z"/>`),
		g.Group(children),
	)
}