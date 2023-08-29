package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderShared(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 8H16l-3.41-3.41A2 2 0 0 0 11.17 4H4a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V10a2 2 0 0 0-2-2Zm-6 18h-8v-1a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1Zm6 0h-4v-1a3 3 0 0 0-3-3h-6a3 3 0 0 0-3 3v1H4V6h7.17l3.42 3.41l.58.59H28Z"/><path fill="currentColor" d="M14 17a4 4 0 1 0 4-4a4 4 0 0 0-4 4Zm4-2a2 2 0 1 1-2 2a2 2 0 0 1 2-2Z"/>`),
		g.Group(children),
	)
}