package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HostMaintenance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m14 23l6-6m1-3a2 2 0 1 0 2 2M17 4h1v1h-1V4Zm-7 19H3V1h18v10M3 13h14M3 18h10M3 8h18"/>`),
		g.Group(children),
	)
}