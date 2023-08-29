package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h6a1 1 0 0 1 .707.293L12.414 5H20a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2ZM4 7v12h16V7H4Zm9 10h-2v-3H8v-2h3V9h2v3h3v2h-3v3Z"/>`),
		g.Group(children),
	)
}