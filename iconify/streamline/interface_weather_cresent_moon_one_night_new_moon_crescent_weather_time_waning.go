package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceWeatherCresentMoonOneNightNewMoonCrescentWeatherTimeWaning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8 7a7 7 0 0 1 3.44-6A6.58 6.58 0 0 0 9 .5a6.5 6.5 0 0 0 0 13a6.58 6.58 0 0 0 2.47-.5A7 7 0 0 1 8 7Z"/>`),
		g.Group(children),
	)
}