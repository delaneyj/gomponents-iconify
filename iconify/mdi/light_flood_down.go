package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightFloodDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 11H2V9h6c.55 0 1 .45 1 1s-.45 1-1 1m6.87-6.65l-.37.92l-6.32-.41L6.9 8H8c1.1 0 2 .9 2 2s-.9 2-2 2H5.27l-.86 2.12l4.8 4.12l-.38.92l1.86.76L16.72 5.1l-1.85-.75m5.94 1.94l.76 1.85l-2.31.96l-.76-1.85l2.31-.96m-2.03 7.28l2.77 1.15l-.76 1.85L18 15.42l.78-1.85m-2.63 7.66L14.3 22l-.96-2.31l1.85-.76l.96 2.3Z"/>`),
		g.Group(children),
	)
}