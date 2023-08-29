package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceWeatherFahrenheitDegreesTemperatureFahrenheitDegreeWeather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="2.5" cy="2" r="1.5"/><path d="M7 13.5v-12h6M7 7h4"/></g>`),
		g.Group(children),
	)
}