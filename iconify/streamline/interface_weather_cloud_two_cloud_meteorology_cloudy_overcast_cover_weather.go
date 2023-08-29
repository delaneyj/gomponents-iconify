package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceWeatherCloudTwoCloudMeteorologyCloudyOvercastCoverWeather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11 6.5a2.51 2.51 0 0 0-1.5.5h0A4.5 4.5 0 1 0 5 11.5h6a2.5 2.5 0 0 0 0-5Z"/>`),
		g.Group(children),
	)
}