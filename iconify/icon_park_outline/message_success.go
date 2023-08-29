package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageSuccess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M25.5 36H21l-10 5v-5H4V6h40v11m-32-3h6m-6 6h12"/><path d="m29 30l6 5l9-11"/></g>`),
		g.Group(children),
	)
}