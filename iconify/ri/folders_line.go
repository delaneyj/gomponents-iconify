package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoldersLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 7V4a1 1 0 0 1 1-1h6.414l2 2H21a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-3v3a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h3Zm0 2H4v10h12v-2H6V9Zm2-4v10h12V7h-5.414l-2-2H8Z"/>`),
		g.Group(children),
	)
}