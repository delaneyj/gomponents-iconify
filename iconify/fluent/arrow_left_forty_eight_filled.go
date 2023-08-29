package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M44.25 24a1.5 1.5 0 0 1-1.5 1.5H10.915l12.387 12.18a1.5 1.5 0 1 1-2.104 2.14L6.201 25.072a.61.61 0 0 1-.02-.02a1.5 1.5 0 0 1 .042-2.145L21.198 8.18a1.5 1.5 0 1 1 2.104 2.14L10.915 22.5H42.75a1.5 1.5 0 0 1 1.5 1.5Z"/>`),
		g.Group(children),
	)
}