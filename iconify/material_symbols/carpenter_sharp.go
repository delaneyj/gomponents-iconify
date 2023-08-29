package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarpenterSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.075 22.7l-4.25-4.225l1.425-1.425L3.1 5.4L7 1.5l14.15 14.125l-7.075 7.075Zm0-2.825L18.3 15.65l-1.4-1.425l-4.25 4.25l1.425 1.4Z"/>`),
		g.Group(children),
	)
}