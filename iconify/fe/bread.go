package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bread(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feBread0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBread1" fill="currentColor" fill-rule="nonzero"><path id="feBread2" d="M17 19v-8.689l.999-.577A1.998 1.998 0 0 0 19 8V7a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v1c0 .723.386 1.377 1.001 1.734L7 10.31V19h10Zm2 0a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-7.535A3.998 3.998 0 0 1 3 8V7a4 4 0 0 1 4-4h10a4 4 0 0 1 4 4v1c0 1.48-.804 2.773-2 3.465V19Z"/></g></g>`),
		g.Group(children),
	)
}