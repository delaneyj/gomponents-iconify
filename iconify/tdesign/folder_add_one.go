package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderAddOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v16H1V2.5Zm2 2V19h18V7H10.52l-2-2.5H3ZM13 9v3h3v2h-3v3h-2v-3H8v-2h3V9h2Z"/>`),
		g.Group(children),
	)
}