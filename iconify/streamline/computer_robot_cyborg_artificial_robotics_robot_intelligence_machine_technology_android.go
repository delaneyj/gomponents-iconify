package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerRobotCyborgArtificialRoboticsRobotIntelligenceMachineTechnologyAndroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.5 9.5c0-3-2.46-5-5.5-5s-5.5 2-5.5 5v4h11ZM7 4.5v-4m-5.5 10h11m-7.5 0v3m4-3v3"/><circle cx="5" cy="8" r=".5"/><circle cx="9" cy="8" r=".5"/></g>`),
		g.Group(children),
	)
}