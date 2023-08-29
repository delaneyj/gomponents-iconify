package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sadapay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.392 44.559L9.525 30.692m13.4-27.133l13.867 13.868m-21.588 8.22l7.624-7.625c4.02-4.02 10.313-4.245 14.055-.503h0c3.743 3.742 3.518 10.035-.503 14.055l-7.666 7.667"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.113 22.472l-7.624 7.625c-4.02 4.02-10.313 4.245-14.055.503h0c-3.743-3.742-3.518-10.035.503-14.055l7.666-7.667"/>`),
		g.Group(children),
	)
}