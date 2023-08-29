package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowExpandLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 22h2V2h-2v9H5.83l5.5-5.5l-1.41-1.42L2 12l7.92 7.92l1.41-1.42l-5.5-5.5H20v9Z"/>`),
		g.Group(children),
	)
}