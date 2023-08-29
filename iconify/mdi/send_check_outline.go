package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendCheckOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 3v7l15 2l-15 2v7l21-9M4 6.03l7.53 3.22l-7.53-1m7.53 6.5L4 17.97v-2.22m18-.25L18.5 19l-2-2l-1.5 1.5l3.5 3.5l5-5Z"/>`),
		g.Group(children),
	)
}