package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceWeatherSunOneCloudMeteorologyCloudyPartlySunnyWeather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 1.93a3 3 0 0 1 4.5 2.57a3.12 3.12 0 0 1-.21 1.11m-2.317 1.56a3.421 3.421 0 0 0-.46 0a3.43 3.43 0 0 0-6.71 0c-.15-.01-.3-.01-.45 0a2.67 2.67 0 1 0 0 5.33h7.62a2.67 2.67 0 0 0 0-5.33v0Z"/>`),
		g.Group(children),
	)
}