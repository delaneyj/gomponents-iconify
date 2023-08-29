package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceWeatherRainDropDropsRainRainyMeteorologyWaterPrecipitationWeather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5 9C11.5 6.51 7 .5 7 .5S2.5 6.51 2.5 9a4.5 4.5 0 0 0 9 0Z"/>`),
		g.Group(children),
	)
}