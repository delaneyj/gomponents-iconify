package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavaidSeaplane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 18a3.996 3.996 0 0 1-3 3.858V17h2v-2h-2v-1.184a3 3 0 1 0-2 0V15h-2v2h2v4.858A3.996 3.996 0 0 1 12 18h-2a6 6 0 0 0 12 0Zm-4-8a1 1 0 1 1-1 1a1 1 0 0 1 1-1Z"/><path fill="currentColor" d="M30 14h-2.18A12.011 12.011 0 0 0 18 4.18V2h-4v2.18A12.011 12.011 0 0 0 4.18 14H2v4h2.18A12.011 12.011 0 0 0 14 27.82V30h4v-2.18A12.011 12.011 0 0 0 27.82 18H30ZM16 26a10 10 0 1 1 10-10a10.011 10.011 0 0 1-10 10Z"/>`),
		g.Group(children),
	)
}