package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndentRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4H2v2h20V4H3Zm6 7H8v2h14v-2H9Zm-7 7h20v2H2v-2Zm3.805-5.293L6.512 12l-.707-.707l-1.768-1.768l-.707-.707l-1.414 1.414l.707.708L3.683 12l-1.06 1.06l-.707.708l1.414 1.414l.707-.707l1.768-1.768Z"/>`),
		g.Group(children),
	)
}