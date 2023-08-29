package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoInputScart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.6 2.2l-3.3.2l-3.5 2l-.5-.9L2 10v7h1v2c0 1.1.9 2 2 2h10c1.1 0 2-.9 2-2v-2h1v-7h-1l-.2-.4l3.5-2l1.8-2.8l-1.5-2.6M15 17v2H5v-2h10Z"/>`),
		g.Group(children),
	)
}