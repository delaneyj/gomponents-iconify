package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeatherMoonTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.36 13.997a7.981 7.981 0 0 1-13.485.541a.599.599 0 0 1 .292-.903c3.006-1.076 4.616-2.323 5.55-4.107c.984-1.877 1.238-3.934.55-6.753a.599.599 0 0 1 .614-.74a7.981 7.981 0 0 1 6.478 11.962Z"/>`),
		g.Group(children),
	)
}