package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceWeatherCloudOneCloudMeteorologyCloudyOvercastCoverWeather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.726 5.706a2.34 2.34 0 0 0-.449 0a3.356 3.356 0 0 0-6.554 0a2.34 2.34 0 0 0-.45 0a2.614 2.614 0 1 0 0 5.219h7.453a2.613 2.613 0 1 0 0-5.219Z"/>`),
		g.Group(children),
	)
}