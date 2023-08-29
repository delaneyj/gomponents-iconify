package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThinLongUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.5 5.5L12 2L8.5 5.5l.707.707L11.5 3.914V22h1V3.914l2.293 2.293l.707-.707Z"/>`),
		g.Group(children),
	)
}