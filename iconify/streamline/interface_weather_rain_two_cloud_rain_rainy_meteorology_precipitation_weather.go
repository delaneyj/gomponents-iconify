package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceWeatherRainTwoCloudRainRainyMeteorologyPrecipitationWeather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4 11.5l.5-1m4 1l.5-1m-3.25 3l.5-1m-5 1l.5-1m8.5 1l.5-1m-.5-4.5a2.5 2.5 0 0 0 0-5a2.54 2.54 0 0 0-1.57.55A3.75 3.75 0 1 0 5 8Z"/>`),
		g.Group(children),
	)
}