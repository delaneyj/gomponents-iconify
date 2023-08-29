package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Digitalwellbeing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="11.06" r="6.47" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15 34.34l12-12.29a6.31 6.31 0 0 1 9.15 0a6.65 6.65 0 0 1 0 9.15C32.05 35.42 24 43.59 24 43.59s-8-8.28-12.07-12.39a6.4 6.4 0 0 1 0-9.15a6.63 6.63 0 0 1 9.15 0l12.06 12.29"/>`),
		g.Group(children),
	)
}