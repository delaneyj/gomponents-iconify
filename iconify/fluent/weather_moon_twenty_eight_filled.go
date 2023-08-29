package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeatherMoonTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.668 2.97a.75.75 0 0 1 .788-.93C20.627 2.536 25.48 7.7 25.48 14c0 6.628-5.372 12-12 12a11.995 11.995 0 0 1-10.378-5.972a.75.75 0 0 1 .469-1.106c1.599-.393 3.55-1.024 5.32-2.004c1.773-.983 3.31-2.287 4.168-4.003c1.714-3.427 1.261-7.345.609-9.945Z"/>`),
		g.Group(children),
	)
}