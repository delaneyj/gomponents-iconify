package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldOffBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2.5"><path stroke-linejoin="round" d="m5.7 5.7l-.094.04A1 1 0 0 0 5 6.66v6.858a6 6 0 0 0 3.023 5.21L12 21l3.977-2.273a5.993 5.993 0 0 0 1.517-1.233M9.66 4.003L12 3l6.394 2.74a1 1 0 0 1 .606.92v6.683"/><path d="m4 4l16 16"/></g>`),
		g.Group(children),
	)
}