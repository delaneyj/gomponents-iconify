package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThinLongLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 8.5L2 12l3.5 3.5l.707-.707L3.914 12.5H22v-1H3.914l2.293-2.293L5.5 8.5Z"/>`),
		g.Group(children),
	)
}