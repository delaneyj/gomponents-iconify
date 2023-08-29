package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mangatoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.741 16.496c0-1.023-.954-1.91-2.054-1.842c-1.027.068-1.833.955-1.833 1.91v1.707c0 1.023.88 1.842 1.98 1.842s1.98-.819 1.98-1.842h-1.98M7.9 20.11v-5.46l2.934 5.46l2.934-5.46v5.46m9.216 0v-5.46l3.888 5.46v-5.46m-6.552 5.46l-1.907-5.46l-1.981 5.46m.66-1.843h2.568M40.1 20.11l-1.907-5.46l-1.98 5.46m.66-1.843h2.567m-26.142 5.032h3.693m-1.812 5.36v-5.36m3.628 3.551c0 1.005.836 1.81 1.812 1.81c1.045 0 1.881-.804 1.881-1.81v-1.809c0-1.005-.836-1.81-1.881-1.81s-1.812.805-1.812 1.81v1.81Zm6.099 0c0 1.005.836 1.81 1.812 1.81c1.045 0 1.881-.804 1.881-1.81v-1.809c0-1.005-.836-1.81-1.881-1.81s-1.812.805-1.812 1.81v1.81Zm6.103 1.809v-5.36l3.693 5.36v-5.36M19.66 31.776c2.352 2.261 5.988 1.929 8.032 0"/>`),
		g.Group(children),
	)
}