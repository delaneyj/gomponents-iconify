package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 3h7v1h1v1h10v1h1v12h-1v1H2v-1H1V4h1V3m1 4v10h16V7H3Z"/>`),
		g.Group(children),
	)
}