package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderKeyholeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.414 3l2 2H21a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h7.414Zm-.828 2H4v14h16V7h-8.414l-2-2ZM12 9a2 2 0 0 1 1.001 3.732L13 17h-2v-4.268A2 2 0 0 1 12 9Z"/>`),
		g.Group(children),
	)
}