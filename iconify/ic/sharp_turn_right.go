package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpTurnRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.17 11l-1.58 1.59L17 14l4-4l-4-4l-1.41 1.41L17.17 9H7v11h2v-9z"/>`),
		g.Group(children),
	)
}