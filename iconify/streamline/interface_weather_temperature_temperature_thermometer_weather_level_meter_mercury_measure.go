package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceWeatherTemperatureTemperatureThermometerWeatherLevelMeterMercuryMeasure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 9V2A1.5 1.5 0 0 0 7 .5h0A1.5 1.5 0 0 0 5.5 2v7a2.5 2.5 0 1 0 3 0Z"/>`),
		g.Group(children),
	)
}