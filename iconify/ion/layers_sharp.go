package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M480 150L256 48L32 150l224 104l224-104zM255.71 392.95l-144.81-66.2L32 362l224 102l224-102l-78.69-35.3l-145.6 66.25z"/><path fill="currentColor" d="m480 256l-75.53-33.53L256.1 290.6l-148.77-68.17L32 256l224 102l224-102Z"/>`),
		g.Group(children),
	)
}