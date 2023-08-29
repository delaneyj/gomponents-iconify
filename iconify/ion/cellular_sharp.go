package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellularSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M496 432h-96V80h96Zm-128 0h-96V160h96Zm-128 0h-96V224h96Zm-128 0H16V288h96Z"/>`),
		g.Group(children),
	)
}