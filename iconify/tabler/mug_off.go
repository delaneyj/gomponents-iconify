package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MugOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5h5.917A1.08 1.08 0 0 1 16 6.077V12m-.167 3.88A4.33 4.33 0 0 1 11.667 19H7.333C4.94 19 3 17.071 3 14.692V6.077A1.08 1.08 0 0 1 4.083 5H5m11 3h2.5c1.38 0 2.5 1.045 2.5 2.333v2.334c0 1.148-.89 2.103-2.06 2.297M3 3l18 18"/>`),
		g.Group(children),
	)
}