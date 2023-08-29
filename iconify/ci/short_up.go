package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShortUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13 7.83l3.59 3.58L18 10l-6-6l-6 6l1.41 1.41L11 7.83V20h2V7.83Z"/>`),
		g.Group(children),
	)
}