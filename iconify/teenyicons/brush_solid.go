package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrushSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 2.5A2.5 2.5 0 0 1 3.5 0H5v5h1V0h1v3h1V0h3.5A2.5 2.5 0 0 1 14 2.5v7a2.5 2.5 0 0 1-2.5 2.5H9v1.5a1.5 1.5 0 0 1-3 0V12H3.5A2.5 2.5 0 0 1 1 9.5v-7ZM2 8v1.5A1.5 1.5 0 0 0 3.5 11H7v2.5a.5.5 0 0 0 1 0V11h3.5A1.5 1.5 0 0 0 13 9.5V8H2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}