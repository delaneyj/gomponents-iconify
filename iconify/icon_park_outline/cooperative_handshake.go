package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CooperativeHandshake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m24 40l12-12l-4 4l-4 4l-4 4Zm0 0L4 20L16 8l8 8"/><path d="M17 23L32 8l12 12l-8 8l-8-8l-6 6l-5-3Zm0 0l7-7m4 20l-3-3m7-1l-3-3"/></g>`),
		g.Group(children),
	)
}