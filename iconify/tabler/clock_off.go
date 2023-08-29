package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.633 5.64a9 9 0 1 0 12.735 12.72m1.674-2.32A9 9 0 0 0 7.96 3.958M12 7v1M3 3l18 18"/>`),
		g.Group(children),
	)
}