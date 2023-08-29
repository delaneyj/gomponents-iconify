package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceWeatherSunTwoCloudMeteorologyCloudyPartlySunnyWeather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11 8.5a2.51 2.51 0 0 0-1.5.5h0A4.5 4.5 0 1 0 5 13.5h6a2.5 2.5 0 0 0 0-5Z"/><circle cx="11.25" cy="2.75" r="2.25"/></g>`),
		g.Group(children),
	)
}