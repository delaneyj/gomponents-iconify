package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartScatter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 6v36h36"/><path fill="currentColor" fill-rule="evenodd" d="M20 24a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm17-8a5 5 0 1 0 0-10a5 5 0 0 0 0 10ZM15 36a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm18-4a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}