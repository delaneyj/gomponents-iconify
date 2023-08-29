package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderBlocked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v5h-2V7H10.52l-2-2.5H3V19h8.5v2H1V2.5Zm17.5 11a3.5 3.5 0 0 0-3.08 5.165l4.745-4.744A3.483 3.483 0 0 0 18.5 13.5Zm3.08 1.835l-4.745 4.744a3.5 3.5 0 0 0 4.745-4.745ZM13 17a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0Z"/>`),
		g.Group(children),
	)
}