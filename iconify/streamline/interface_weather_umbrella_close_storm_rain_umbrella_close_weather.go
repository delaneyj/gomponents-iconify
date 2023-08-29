package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceWeatherUmbrellaCloseStormRainUmbrellaCloseWeather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 9.5h-7l3-9h1l3 9zM7 9.5V12a1.5 1.5 0 0 0 3 0"/>`),
		g.Group(children),
	)
}