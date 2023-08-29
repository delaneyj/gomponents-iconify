package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlertMinusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 15.5h2v2h-2v-2m3 3.5v-.4H5.4L12 7.3l4.11 7.14c.51-.44 1.09-.79 1.73-1.03L12 3.3L2 20.6h12.22c-.14-.51-.22-1.04-.22-1.6m-1-8.5h-2v4h2v-4m3 7.5v2h8v-2h-8Z"/>`),
		g.Group(children),
	)
}