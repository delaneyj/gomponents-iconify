package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneRing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M506 316q6 6 6 15t-6 15l-53 53q-6 6-15 6t-15-6q-26-24-57-40q-12-5-12-19v-66q-47-15-98-15t-98 15v66q0 14-12 20q-32 16-57 39q-6 6-15 6t-15-6L6 346q-6-6-6-15t6-15q105-100 250-100t250 100zM451 94l-76 75l-30-30l76-76zM277 3v106h-42V3h42zM137 169Q63 94 61 94l30-31l76 76z"/>`),
		g.Group(children),
	)
}