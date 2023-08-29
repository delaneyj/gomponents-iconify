package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThinBigUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><g clip-path="url(#ciThinBigUp0)"><path fill="currentColor" d="M19.5 6.5L13 0L6.5 6.5l.707.707L12.5 1.914V24h1V1.914l5.293 5.293l.707-.707Z"/></g><defs><clipPath id="ciThinBigUp0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}