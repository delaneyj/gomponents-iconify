package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileChartCheckOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m23.5 17l-5 5l-3.5-3.5l1.5-1.5l2 2l3.5-3.5l1.5 1.5M6 2c-1.1 0-2 .9-2 2v16c0 1.1.9 2 2 2h7.8c-.4-.6-.6-1.3-.7-2H6V4h7v5h5v4.1c.3-.1.7-.1 1-.1c.3 0 .7 0 1 .1V8l-6-6H6m5 9v8h2v-8h-2m-4 2v6h2v-6H7Z"/>`),
		g.Group(children),
	)
}