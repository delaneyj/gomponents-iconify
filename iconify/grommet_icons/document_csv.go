package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentCsv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 9V1H19.5L23 4.5V23H4M18 1v5h5M7 13H5c-1 0-2 .5-2 1.5v3c0 1 1 1.5 2 1.5h2m6.25-6h-2.5c-1.5 0-2 .5-2 1.5s.5 1.5 2 1.5s2 .5 2 1.5s-.5 1.5-2 1.5h-2.5m12.25-7v.5C20.5 13 18 19 18 19h-.5S15 13 15 12.5V12"/>`),
		g.Group(children),
	)
}