package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M3 10V7h.5a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5H3Zm4-2.5a.5.5 0 0 1 1 0v2a.5.5 0 0 1-1 0v-2Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM3.5 6H2v5h1.5A1.5 1.5 0 0 0 5 9.5v-2A1.5 1.5 0 0 0 3.5 6Zm4 0A1.5 1.5 0 0 0 6 7.5v2a1.5 1.5 0 0 0 3 0v-2A1.5 1.5 0 0 0 7.5 6Zm2.5 5V6h3v2h-1V7h-1v3h1V9h1v2h-3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}