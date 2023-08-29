package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloorLampTorchiere(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7 2l3 7h4l3-7H7m6 18h3v2H8v-2h3V10h2v10Z"/>`),
		g.Group(children),
	)
}