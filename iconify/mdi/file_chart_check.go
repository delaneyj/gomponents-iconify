package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileChartCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m23.5 17l-5 5l-3.5-3.5l1.5-1.5l2 2l3.5-3.5l1.5 1.5M6 2c-1.1 0-2 .9-2 2v16c0 1.1.9 2 2 2h7.8c-.5-.9-.8-1.9-.8-3v1h-2v-8h2v7c0-3.3 2.7-6 6-6c.3 0 .7 0 1 .1V8l-6-6H6m7 1.5L18.5 9H13V3.5M7 14h2v6H7v-6Z"/>`),
		g.Group(children),
	)
}