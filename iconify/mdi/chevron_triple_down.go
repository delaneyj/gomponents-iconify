package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronTripleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.41 14.58L12 19.17l4.59-4.59L18 16l-6 6l-6-6l1.41-1.42m0-6L12 13.17l4.59-4.59L18 10l-6 6l-6-6l1.41-1.42m0-6L12 7.17l4.59-4.59L18 4l-6 6l-6-6l1.41-1.42Z"/>`),
		g.Group(children),
	)
}