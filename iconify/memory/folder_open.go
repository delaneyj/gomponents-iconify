package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M1 4h1V3h7v1h1v1h10v1h1v12h-1v1H2v-1H1V4m2 5h16V7H9V6H8V5H3v4m0 8h16v-6H3v6Z"/>`),
		g.Group(children),
	)
}