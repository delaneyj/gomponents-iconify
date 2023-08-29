package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartHistogram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 3v18h18m-1-3v3m-4-5v5m-4-8v8m-4-5v5"/><path d="M3 11c6 0 5-5 9-5s3 5 9 5"/></g>`),
		g.Group(children),
	)
}