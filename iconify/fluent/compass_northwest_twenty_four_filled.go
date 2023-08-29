package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassNorthwestTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm.458-12.614a4 4 0 0 1 2.196 2.14l1.854 4.325a.5.5 0 0 1-.657.657l-4.325-1.854a4 4 0 0 1-2.14-2.196L7.78 8.43a.5.5 0 0 1 .65-.65l4.028 1.606Z"/>`),
		g.Group(children),
	)
}