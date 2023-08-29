package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeatherSunnyTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 2a.75.75 0 0 1 .75.75v1.496a.75.75 0 0 1-1.5 0V2.75A.75.75 0 0 1 14 2Zm6 12a6 6 0 1 1-12 0a6 6 0 0 1 12 0Zm5.25.75a.75.75 0 0 0 0-1.5h-1.496a.75.75 0 0 0 0 1.5h1.496ZM14 23.004a.75.75 0 0 1 .75.75v1.496a.75.75 0 0 1-1.5 0v-1.496a.75.75 0 0 1 .75-.75ZM4.25 14.75a.75.75 0 0 0 0-1.5H2.751a.75.75 0 0 0 0 1.5h1.5Zm.97-9.531a.75.75 0 0 1 1.06 0l1.5 1.5a.75.75 0 0 1-1.06 1.06l-1.5-1.5a.75.75 0 0 1 0-1.06Zm1.06 17.56a.75.75 0 0 1-1.06-1.06l1.5-1.5a.75.75 0 0 1 1.06 1.06l-1.5 1.5Zm16.5-17.56a.75.75 0 0 0-1.06 0l-1.5 1.5a.75.75 0 1 0 1.06 1.06l1.5-1.5a.75.75 0 0 0 0-1.06ZM21.72 22.78a.75.75 0 0 0 1.06-1.06l-1.5-1.5a.75.75 0 0 0-1.06 1.06l1.5 1.5Z"/>`),
		g.Group(children),
	)
}