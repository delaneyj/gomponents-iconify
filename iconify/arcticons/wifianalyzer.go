package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wifianalyzer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 23.87A24.38 24.38 0 0 1 24 14h0a24.38 24.38 0 0 1 19.5 9.85m-32 5.23A15.69 15.69 0 0 1 24 22.55a15.69 15.69 0 0 1 12.53 6.55M18 34a7.52 7.52 0 0 1 12 0"/>`),
		g.Group(children),
	)
}