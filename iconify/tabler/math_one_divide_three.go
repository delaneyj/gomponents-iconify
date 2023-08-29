package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MathOneDivideThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 15.5a.5.5 0 0 1 .5-.5h2a1.5 1.5 0 0 1 0 3h-1.167H12.5a1.5 1.5 0 0 1 0 3h-2a.5.5 0 0 1-.5-.5M5 12h14m-9-7l2-2v6"/>`),
		g.Group(children),
	)
}