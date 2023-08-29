package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thermoworks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.608 38.5a15.054 15.054 0 0 0 1.313-6.601c0-7.894-3.842-7.894-3.842-15.788A15.072 15.072 0 0 1 23.395 9.5m8.792 29a15.054 15.054 0 0 0 1.313-6.601c0-7.894-3.842-7.894-3.842-15.787A15.072 15.072 0 0 1 30.974 9.5m-13.945 29a15.054 15.054 0 0 0 1.313-6.601c0-7.894-3.842-7.894-3.842-15.788A15.072 15.072 0 0 1 15.815 9.5"/>`),
		g.Group(children),
	)
}