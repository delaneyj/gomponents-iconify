package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeTemperatureOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 8L4.311 3.156a.6.6 0 0 0-.6.037L2.5 4m9.5 7v8a2 2 0 0 1-2 2H7m0 0H3.6a.6.6 0 0 1-.6-.6v-4.8a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6V21Zm12-3a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0-10.5V14m0-2h2m-2-3h2"/>`),
		g.Group(children),
	)
}