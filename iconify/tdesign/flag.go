package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 2h8.414l2 2h8.014l-2.364 6.5l2.364 6.5h-8.842l-2-2H5v7.5H3V2Zm2 11h6.414l2 2h5.158l-1.636-4.5L18.572 6h-5.986l-2-2H5v9Z"/>`),
		g.Group(children),
	)
}