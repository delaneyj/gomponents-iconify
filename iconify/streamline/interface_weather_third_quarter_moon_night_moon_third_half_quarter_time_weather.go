package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceWeatherThirdQuarterMoonNightMoonThirdHalfQuarterTimeWeather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.25 13.5a6.5 6.5 0 0 1 0-13Z"/>`),
		g.Group(children),
	)
}