package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dashboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm3-6a3 3 0 0 0-6 0M5 5l2 2m5 0v6m0-10v2m7 7h2M3 12h2m12-5l2-2M3 17h18"/>`),
		g.Group(children),
	)
}