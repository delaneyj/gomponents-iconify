package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpTurnSharpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18 6.83l1.59 1.58L21 7l-4-4l-4 4l1.41 1.41L16 6.83V13H6v8h2v-6h10z"/>`),
		g.Group(children),
	)
}