package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 2H5C3.3 2 2 3.3 2 5v14c0 1.7 1.3 3 3 3h14c1.7 0 3-1.3 3-3V5c0-1.7-1.3-3-3-3zm1 11.9L18.1 12c-1.2-1.1-3.1-1.1-4.2 0l-.9.9l-2.9-2.9C8.9 8.9 7 8.9 5.9 10L4 11.9V5c0-.6.4-1 1-1h14c.6 0 1 .4 1 1v8.9z"/>`),
		g.Group(children),
	)
}