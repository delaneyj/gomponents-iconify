package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartRadar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 3l9.5 7L18 21H6L2.5 10z"/><path d="m12 7.5l5.5 4L15 17H8.5l-2-5.5z"/><path d="m2.5 10l9.5 3l9.5-3"/><path d="M12 3v10l6 8M6 21l6-8"/></g>`),
		g.Group(children),
	)
}