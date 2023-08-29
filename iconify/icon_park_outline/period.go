package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Period(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="square" stroke-width="4" d="M8 4c4.002 3.337 6.003 6.671 6.003 10.005c0 5-4.003 7.239-6.003 9.995c-2 2.756-2.995 5.931-2.995 10.003c0 2.715.998 6.047 2.995 9.997M40.003 4C36.001 7.337 34 10.671 34 14.005c0 5 4.004 7.239 6.004 9.995c2 2.756 2.995 5.931 2.995 10.003c0 2.715-.998 6.047-2.995 9.997m-26-14.817c2.522-.242 4.615.118 6.281 1.08c2.5 1.445 3.707 3.097 3.715 4.694c.006 1.065.006 4.08 0 9.043m10.004-14.817c-2.522-.242-4.615.118-6.281 1.08c-2.5 1.445-3.707 3.097-3.715 4.694"/>`),
		g.Group(children),
	)
}