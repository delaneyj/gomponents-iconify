package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 23.3L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4ZM5 6.425l2 2V18h9.6l2.4 2.4V23H5V6.425ZM19 1v15.15l-2-2V6H8.825L5.15 2.3V1H19Z"/>`),
		g.Group(children),
	)
}