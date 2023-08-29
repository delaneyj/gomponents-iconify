package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaStepBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M402.8 90.58a23.876 23.876 0 0 0-25.082 2.149L183.155 236.847a24 24 0 0 0 0 38.57l194.56 144.119A24 24 0 0 0 416 400.251V112.015a23.882 23.882 0 0 0-13.2-21.435ZM384 384.37L210.881 256.133L384 127.9ZM96 88h32v336H96z"/>`),
		g.Group(children),
	)
}