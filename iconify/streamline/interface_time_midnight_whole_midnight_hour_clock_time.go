package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTimeMidnightWholeMidnightHourClockTime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2 7H.5m2.96-3.54L2.4 2.4m1.06 8.14L2.4 11.6M12 7h1.5m-2.96 3.54l1.06 1.06M7 12v1.5m3.54-10.04L11.6 2.4M7 2V.5"/>`),
		g.Group(children),
	)
}