package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Connectivity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 8v4l3 3l6-6l-4-4H5L2.5 3M17 7l3 3v7m-7-6l3 3m-5-1l3 3m-7 1l3 3h10"/>`),
		g.Group(children),
	)
}