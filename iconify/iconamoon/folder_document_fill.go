package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderDocumentFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 5a1 1 0 0 1 1-1h7a1 1 0 0 1 .707.293L12.414 6H21a1 1 0 0 1 1 1v10a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V5Zm5 5a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2H7Zm0 4a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2H7Zm5-3a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-4Zm2 1v2h2v-2h-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}