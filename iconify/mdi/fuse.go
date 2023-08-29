package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fuse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 7v10h7V7H8m3.16 9v-3.13H9.41L11.91 8v3.14h1.68L11.16 16M16 2v4H7V2a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1m0 16v4a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1v-4h9Z"/>`),
		g.Group(children),
	)
}