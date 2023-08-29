package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronsRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.296 7.71L14.621 12l-4.325 4.29l1.408 1.42L17.461 12l-5.757-5.71z"/><path fill="currentColor" d="M6.704 6.29L5.296 7.71L9.621 12l-4.325 4.29l1.408 1.42L12.461 12z"/>`),
		g.Group(children),
	)
}