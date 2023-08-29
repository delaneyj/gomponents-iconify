package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VibrateLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M160 34H96a22 22 0 0 0-22 22v144a22 22 0 0 0 22 22h64a22 22 0 0 0 22-22V56a22 22 0 0 0-22-22Zm10 166a10 10 0 0 1-10 10H96a10 10 0 0 1-10-10V56a10 10 0 0 1 10-10h64a10 10 0 0 1 10 10Zm44-112v80a6 6 0 0 1-12 0V88a6 6 0 0 1 12 0Zm32 16v48a6 6 0 0 1-12 0v-48a6 6 0 0 1 12 0ZM54 88v80a6 6 0 0 1-12 0V88a6 6 0 0 1 12 0Zm-32 16v48a6 6 0 0 1-12 0v-48a6 6 0 0 1 12 0Z"/>`),
		g.Group(children),
	)
}