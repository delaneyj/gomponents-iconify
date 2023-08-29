package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M.09 1C.03 1 0 1.04 0 1.09V6.9c0 .05.04.09.09.09H6.9c.05 0 .09-.04.09-.09V4.99h1v-2h-1V1.08c0-.06-.04-.09-.09-.09H.09zM1 2h5v4H1V2z"/>`),
		g.Group(children),
	)
}