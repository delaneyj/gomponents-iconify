package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Conekta(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M11 26.005L3.599 16L11 6l4.599 6.198L12.594 16l3 4l-4.599 6.005zM16.599 32l11.802-16L16.599 0H7.401l12 16l-12 16z"/>`),
		g.Group(children),
	)
}