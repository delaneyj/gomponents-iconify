package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ohmsnow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.5 29l5.5 8.36h10.14C13.57 24.54 14.68 10.62 24 10.64s10.3 13.88 3.88 26.72H38L43.5 29"/>`),
		g.Group(children),
	)
}