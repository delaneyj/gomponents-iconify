package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceWeatherWindWindHighOvercastGustWeatherMeteorologyGale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5.5a1.75 1.75 0 0 1 0 3.5h-7m11.25 6.5a1.75 1.75 0 0 0 0-3.5H2m5.25 6.5a1.75 1.75 0 0 0 0-3.5H1.5"/>`),
		g.Group(children),
	)
}