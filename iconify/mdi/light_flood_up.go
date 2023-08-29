package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightFloodUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 15H2v-2h6c.55 0 1 .45 1 1s-.45 1-1 1m8.72 3.9L10.69 4.08l-1.86.76l.38.92l-4.8 4.12l.86 2.12H8c1.1 0 2 .9 2 2s-.9 2-2 2H6.9l1.28 3.14l6.32-.41l.37.92l1.85-.75m4.09-1.19l-2.31-.96l.76-1.85l2.31.96l-.76 1.85m-2.03-7.28L18 8.58l2.79-1.15l.76 1.85l-2.77 1.15m-3.59-5.36l-1.85-.76L14.3 2l1.85.77l-.96 2.3Z"/>`),
		g.Group(children),
	)
}