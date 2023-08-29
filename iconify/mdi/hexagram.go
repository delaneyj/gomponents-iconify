package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hexagram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.5 7.66l6.14-.69L17 12l3.68 4.97l-6.18-.65L12.03 22L9.5 16.34l-6.14.69L7 12L3.32 7.03l6.18.65L11.97 2l2.53 5.66Z"/>`),
		g.Group(children),
	)
}