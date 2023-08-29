package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeatherSnowflakeTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 2a1 1 0 0 1 1 1v2.551l1.707-1.706a1 1 0 1 1 1.414 1.414L13 8.379V11h2.62l3.121-3.12a1 1 0 1 1 1.414 1.413L18.45 11H21a1 1 0 1 1 0 2h-2.551l1.706 1.707a1 1 0 0 1-1.414 1.414L15.621 13H13v2.62l3.12 3.121a1 1 0 0 1-1.413 1.414L13 18.45V21a1 1 0 1 1-2 0v-2.551l-1.707 1.706a1 1 0 0 1-1.414-1.414L11 15.621V13H8.38l-3.121 3.12a1 1 0 0 1-1.414-1.413L5.55 13H3a1 1 0 1 1 0-2h2.551L3.845 9.293A1 1 0 1 1 5.259 7.88L8.379 11H11V8.38L7.88 5.259a1 1 0 0 1 1.413-1.414L11 5.55V3a1 1 0 0 1 1-1Z"/>`),
		g.Group(children),
	)
}