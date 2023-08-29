package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneMissed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M139 53v75h-32V0h128v32h-75l96 96L384 0l21 21l-149 150zm367 239q6 6 6 15t-6 15l-53 53q-6 6-15 6t-15-6q-27-24-57-40q-12-5-12-19v-66q-47-15-98-15t-98 15v66q0 14-12 20q-32 16-57 39q-6 6-15 6t-15-6L6 322q-6-6-6-15t6-15q105-100 250-100t250 100z"/>`),
		g.Group(children),
	)
}