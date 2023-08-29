package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 2.5h10.48l2 2.5H23v2H12.52l-2-2.5H1v-2Zm0 4h8.48l2 2.5H23v12H1V6.5Zm2 2V19h18v-8H10.52l-2-2.5H3Z"/>`),
		g.Group(children),
	)
}