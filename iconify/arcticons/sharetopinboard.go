package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sharetopinboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.826 4.5A1.153 1.153 0 0 0 15.8 5.525V6.55a1.078 1.078 0 0 0 1.025.882h3.17c.075 8.414-.59 13.855-4.015 14.073a4.007 4.007 0 0 0-4.004 4.003h10.021v15.99L24 43.5l2.002-2.002v-15.99h10.02a4.007 4.007 0 0 0-4.002-4.003c-3.943-.136-4.333-5.442-4.016-14.073h3.171a1.078 1.078 0 0 0 1.024-.882V5.525A1.153 1.153 0 0 0 31.175 4.5Zm5.172 21.008h4.004"/>`),
		g.Group(children),
	)
}