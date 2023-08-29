package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M10 17v27h28V17"/><path d="m5 22l5-5L24 4l14 13l5 5"/><path d="m19 19l5 6l5-6M18 31h12m-12-6h12m-6 0v12"/></g>`),
		g.Group(children),
	)
}