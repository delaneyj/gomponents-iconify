package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceWeatherHurricaneTwisterTornadoHurricaneCycloneDisasterNaturalWeather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><ellipse cx="7" cy="3" rx="6" ry="2.5"/><path d="M11.87 6.73c-.87.79-2.73 1.34-4.94 1.34A10.17 10.17 0 0 1 3 7.37m7.86 2.47a6.84 6.84 0 0 1-3.36.73a8.47 8.47 0 0 1-2.5-.35m5.37 3.22c-.95.21-1.83-.16-2-.83"/></g>`),
		g.Group(children),
	)
}