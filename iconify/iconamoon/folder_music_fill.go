package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderMusicFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 5a1 1 0 0 1 1-1h7a1 1 0 0 1 .707.293L12.414 6H21a1 1 0 0 1 1 1v10a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V5Zm9 8v-3a.997.997 0 0 1 1-1a.997.997 0 0 1 .707.293l2 2a1 1 0 0 1-1.414 1.414L13 12.414V15a2 2 0 1 1-2-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}