package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M208 55.27v145.46a7.28 7.28 0 0 1-7.27 7.27H55.27a7.28 7.28 0 0 1-7.27-7.27V55.27A7.28 7.28 0 0 1 55.27 48h145.46a7.28 7.28 0 0 1 7.27 7.27Z" opacity=".2"/><path d="M200.73 40H55.27A15.29 15.29 0 0 0 40 55.27v145.46A15.29 15.29 0 0 0 55.27 216h145.46A15.29 15.29 0 0 0 216 200.73V55.27A15.29 15.29 0 0 0 200.73 40ZM200 200H56V56h144Z"/></g>`),
		g.Group(children),
	)
}