package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceWeatherRainOneCloudRainRainyMeteorologyPrecipitationWeather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.25 11.5l.5-1m4 1l.5-1M6 13.5l.5-1m-5 1l.5-1m8.5 1l.5-1M11 8a2.5 2.5 0 1 0-.62-4.92a3.5 3.5 0 0 0-6.76 0A2.5 2.5 0 1 0 3 8Z"/>`),
		g.Group(children),
	)
}