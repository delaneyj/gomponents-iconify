package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TemperatureOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 10v3.5a4 4 0 1 0 5.836 2.33M14 10V5a2 2 0 1 0-4 0v1m3 3h1M3 3l18 18"/>`),
		g.Group(children),
	)
}