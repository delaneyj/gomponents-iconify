package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneEnd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M256 128q-51 0-98 15v66q0 14-12 20q-31 15-57 39q-6 6-15 6t-15-6L6 215q-6-6-6-15t6-15Q111 85 256 85t250 100q6 6 6 15t-6 15l-53 53q-6 6-15 6t-15-6q-25-23-57-39q-12-6-12-19v-66q-47-16-98-16z"/>`),
		g.Group(children),
	)
}