package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneTurnLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.83 11l1.59 1.59L7 14l-4-4l4-4l1.41 1.41L6.83 9H15c1.1 0 2 .9 2 2v9h-2v-9H6.83z"/>`),
		g.Group(children),
	)
}