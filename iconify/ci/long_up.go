package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LongUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13 5.83l2.59 2.58L17 7l-5-5l-5 5l1.41 1.41L11 5.83V22h2V5.83Z"/>`),
		g.Group(children),
	)
}