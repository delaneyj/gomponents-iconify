package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropAltRotateLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.71 16.29a1 1 0 0 0-1.42 1.42l.3.29H8a3 3 0 0 1-3-3v-2a1 1 0 0 0-2 0v2a5 5 0 0 0 5 5h.59l-.3.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l2-2a1 1 0 0 0 0-1.42ZM20 12h-1V5a1 1 0 0 0-1-1h-7V3a1 1 0 0 0-2 0v1H8a1 1 0 0 0 0 2h1v7a1 1 0 0 0 1 1h7v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2Zm-3 0h-6V6h6Z"/>`),
		g.Group(children),
	)
}