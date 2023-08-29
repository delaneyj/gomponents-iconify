package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceWeatherGibbousMoonOneNightMoonWeatherGibbousTimeWaning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2 7a6.5 6.5 0 0 0 6.5 6.5a6.58 6.58 0 0 0 1.19-.11a10 10 0 0 0 0-12.78A6.58 6.58 0 0 0 8.5.5A6.5 6.5 0 0 0 2 7Z"/>`),
		g.Group(children),
	)
}