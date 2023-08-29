package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M5 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm7.5 7H10L9 7L8 5.25L9 5l2.3 1a.539.539 0 0 0 .4-1L9 4H7L5 5L4 6H2.5a.5.5 0 0 0 0 1H5l1-1l1 2l-2 2v3.5a.5.5 0 0 0 1 0v-3.11L8 9l1 2h3.5a.5.5 0 0 0 0-1z"/>`),
		g.Group(children),
	)
}