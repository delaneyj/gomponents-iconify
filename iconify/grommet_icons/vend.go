package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M3 2.002A.998.998 0 0 1 3.993 1h16.014c.548 0 .993.44.993 1.002V23H3V2.002ZM15 1v22M3 16h12m-7 1h2m-3-1v-4m0-3V5m4 11v-4m0-3V5m6 7h1m-1-3h2m-2 11h2M3 9h12"/>`),
		g.Group(children),
	)
}