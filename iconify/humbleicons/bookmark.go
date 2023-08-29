package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M6 20V5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v15l-6-3l-6 3z"/>`),
		g.Group(children),
	)
}