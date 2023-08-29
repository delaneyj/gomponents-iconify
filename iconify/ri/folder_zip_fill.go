package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderZipFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 5a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h7.414l2 2H16v2h2V5h3Zm-3 8h-2v2h-2v3h4v-5Zm-2-2h-2v2h2v-2Zm2-2h-2v2h2V9Zm-2-2h-2v2h2V7Z"/>`),
		g.Group(children),
	)
}