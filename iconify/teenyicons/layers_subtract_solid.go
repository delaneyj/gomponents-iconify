package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersSubtractSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 4H1.5A1.5 1.5 0 0 0 0 5.5v8A1.5 1.5 0 0 0 1.5 15h8a1.5 1.5 0 0 0 1.5-1.5V11h2.5A1.5 1.5 0 0 0 15 9.5v-8A1.5 1.5 0 0 0 13.5 0h-8A1.5 1.5 0 0 0 4 1.5V4Zm1-2.5a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5h-8a.5.5 0 0 1-.5-.5v-8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}