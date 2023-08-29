package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Translation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m17 32l2.188-5M31 32l-2.188-5m-9.625 0L24 16l4.813 11m-9.625 0h9.625"/><path d="M43.2 20c-1.853-9.129-9.924-16-19.6-16C13.924 4 5.853 10.871 4 20l6-2M4 28c1.853 9.129 9.924 16 19.6 16c9.676 0 17.747-6.871 19.6-16L38 30"/></g>`),
		g.Group(children),
	)
}