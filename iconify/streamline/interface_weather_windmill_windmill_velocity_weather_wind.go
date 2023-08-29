package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceWeatherWindmillWindmillVelocityWeatherWind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 .5h1a2 2 0 0 1 2 2V7h0h-3h0V.5h0ZM4 7h3v6.5h0h-1a2 2 0 0 1-2-2V7h0Zm9.5 0v1a2 2 0 0 1-2 2H7h0V7h6.5ZM7 4v3h0H.5h0V6a2 2 0 0 1 2-2H7Z"/>`),
		g.Group(children),
	)
}