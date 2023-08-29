package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipVerticalSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7 0v15h1V0H7ZM4.615 3.013A.5.5 0 0 1 5 3.5v8a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.447-.724l4-8a.5.5 0 0 1 .562-.263Zm5.77 0a.5.5 0 0 1 .562.263l4 8A.5.5 0 0 1 14.5 12h-4a.5.5 0 0 1-.5-.5v-8a.5.5 0 0 1 .385-.487Z"/>`),
		g.Group(children),
	)
}