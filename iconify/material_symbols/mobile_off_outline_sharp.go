package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileOffOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 23.3L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4ZM5 6.425l2 2V18h9.6l2.4 2.4V23H5V6.425ZM7 20v1h10v-1H7ZM19 1v15.15l-2-2V6H8.825L5.15 2.3V1H19ZM7 4h10V3H7v1Zm0 16v1v-1ZM7 4V3v1Z"/>`),
		g.Group(children),
	)
}