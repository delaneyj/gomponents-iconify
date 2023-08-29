package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Octagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.77 1.095a.6.6 0 0 1 .46 0l7.319 3.032a.6.6 0 0 1 .324.324l3.032 7.32a.6.6 0 0 1 0 .459l-3.032 7.319a.6.6 0 0 1-.324.324l-7.32 3.032a.6.6 0 0 1-.459 0l-7.319-3.032a.6.6 0 0 1-.324-.324l-3.032-7.32a.6.6 0 0 1 0-.459l3.032-7.319a.6.6 0 0 1 .324-.324l7.32-3.032Z"/>`),
		g.Group(children),
	)
}