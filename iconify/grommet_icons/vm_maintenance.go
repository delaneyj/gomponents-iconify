package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VmMaintenance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M19 10V2H7v12h7V7H2v12h8m4 4l6-6m1-3a2 2 0 1 0 2 2"/>`),
		g.Group(children),
	)
}