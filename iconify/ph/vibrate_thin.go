package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VibrateThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M160 36H96a20 20 0 0 0-20 20v144a20 20 0 0 0 20 20h64a20 20 0 0 0 20-20V56a20 20 0 0 0-20-20Zm12 164a12 12 0 0 1-12 12H96a12 12 0 0 1-12-12V56a12 12 0 0 1 12-12h64a12 12 0 0 1 12 12Zm40-112v80a4 4 0 0 1-8 0V88a4 4 0 0 1 8 0Zm32 16v48a4 4 0 0 1-8 0v-48a4 4 0 0 1 8 0ZM52 88v80a4 4 0 0 1-8 0V88a4 4 0 0 1 8 0Zm-32 16v48a4 4 0 0 1-8 0v-48a4 4 0 0 1 8 0Z"/>`),
		g.Group(children),
	)
}