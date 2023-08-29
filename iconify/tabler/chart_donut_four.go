package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartDonutFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8.848 14.667L5.5 17.5M12 3v5m4 4h5M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0m11.219 3.328L17 19.5"/><path d="M8 12a4 4 0 1 0 8 0a4 4 0 1 0-8 0"/></g>`),
		g.Group(children),
	)
}