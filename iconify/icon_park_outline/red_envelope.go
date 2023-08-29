package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedEnvelope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M39 4H9v6l15 4l15-4V4Zm0 13v27H9V17"/><path d="m19 19l5 6l5-6M18 31h12m-12-6h12m-6 0v12"/></g>`),
		g.Group(children),
	)
}