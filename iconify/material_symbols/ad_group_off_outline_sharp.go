package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdGroupOffOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.7 17.875L18.825 16H20V6H8.825l-2.7-2.7V2H22v15.875h-1.3Zm-.2 5.425L15.2 18H6V8.8L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4ZM8 16h5.2L8 10.8V16Zm-6 6V6h2v14h14v2H2Zm8.625-8.575Zm2.85-2.775Z"/>`),
		g.Group(children),
	)
}