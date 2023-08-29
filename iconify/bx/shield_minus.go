package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.41 6.11l-8-4a.94.94 0 0 0-.89 0l-8 4A1 1 0 0 0 3 6.9c0 .11-1 10.77 8.59 15a1 1 0 0 0 .41.1a.93.93 0 0 0 .4-.09C21.92 17.67 21 7 21 6.9a1 1 0 0 0-.59-.79zM12 19.9C5.2 16.63 4.88 9.64 4.93 7.63l7-3.51l7 3.51c.13 2.01-.19 9-6.93 12.27z"/><path fill="currentColor" d="M8 11h8v2H8z"/>`),
		g.Group(children),
	)
}