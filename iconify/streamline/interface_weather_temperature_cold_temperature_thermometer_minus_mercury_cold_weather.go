package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceWeatherTemperatureColdTemperatureThermometerMinusMercuryColdWeather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5 9V2A1.5 1.5 0 0 0 10 .5h0A1.5 1.5 0 0 0 8.5 2v7a2.5 2.5 0 1 0 3 0Zm-10-6.5h3.25"/>`),
		g.Group(children),
	)
}