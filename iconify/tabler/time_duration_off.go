package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimeDurationOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12v.01m4.5 7.79v.01M4.2 16.5v.01m0-9.01v.01M12 21a8.994 8.994 0 0 0 6.362-2.634m1.685-2.336A9 9 0 0 0 12 3M3 3l18 18"/>`),
		g.Group(children),
	)
}