package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTimeSixHourClockTimeMinutesHalfThirtySix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 13.5a6.5 6.5 0 0 0 0-13ZM2 7H.5m2.96-3.54L2.4 2.4m1.06 8.14L2.4 11.6"/>`),
		g.Group(children),
	)
}