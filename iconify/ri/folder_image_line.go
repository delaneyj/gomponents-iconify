package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderImageLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 5v14h16V7h-8.414l-2-2H4Zm8.414 0H21a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h7.414l2 2ZM10 10.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm8 6.5l-4-6l-7 6h11Z"/>`),
		g.Group(children),
	)
}