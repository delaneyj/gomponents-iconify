package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatabaseNetworkPoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 36v-6m-4 10H6m22 0h14m-14 0a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm9-23c0 7.18-5.82 13-13 13s-13-5.82-13-13m26 0c0-7.18-5.82-13-13-13S11 9.82 11 17m26 0H11"/><path d="M29 17c0 7.18-2.239 13-5 13s-5-5.82-5-13s2.239-13 5-13s5 5.82 5 13Zm8 0H11"/></g>`),
		g.Group(children),
	)
}