package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceWeatherCelsiusDegreesTemperatureCentigradeCelsiusDegreeWeather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="2" cy="2" r="1.5"/><path d="M13.5 1.84a6 6 0 1 0-2 11.66a6 6 0 0 0 2-.34"/></g>`),
		g.Group(children),
	)
}