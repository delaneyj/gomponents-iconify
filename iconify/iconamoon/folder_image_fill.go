package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderImageFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 5a1 1 0 0 1 1-1h7a1 1 0 0 1 .707.293L12.414 6H21a1 1 0 0 1 1 1v10a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V5Zm4 4.5A1.5 1.5 0 0 1 7.5 8h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H7.5A1.5 1.5 0 0 1 6 9.51V9.5Zm6.707 1.793a1 1 0 0 0-1.414 0l-4 4a1 1 0 1 0 1.414 1.414L12 13.414l3.293 3.293a1 1 0 0 0 1.414-1.414l-4-4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}