package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vial(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 22a4.965 4.965 0 0 0 3.535-1.465l9.193-9.193l.707.708l1.414-1.414l-8.485-8.486l-1.414 1.414l.708.707l-9.193 9.193C2.521 14.408 2 15.664 2 17s.521 2.592 1.465 3.535A4.965 4.965 0 0 0 7 22zM18.314 9.928L15.242 13H6.758l7.314-7.314l4.242 4.242z"/>`),
		g.Group(children),
	)
}