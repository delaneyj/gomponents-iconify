package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderZip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v16H1V2.5Zm2 2V19h11v-2h2v-1.996h-2V13h2v-1.996h-2V9h2V7h-5.48l-2-2.5H3Zm13.004 4.504V11h2v2.004h-2V15h2v2.004h-2V19H21V7h-2.996v2.004h-2Z"/>`),
		g.Group(children),
	)
}