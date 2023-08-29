package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkMinusF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 .858h8a3 3 0 0 1 3 3v15a2 2 0 0 1-3.348 1.477l-2.978-2.717a1 1 0 0 0-1.348 0l-2.978 2.717A2 2 0 0 1 0 18.858v-15a3 3 0 0 1 3-3zm1 8a1 1 0 1 0 0 2h6a1 1 0 0 0 0-2H4z"/>`),
		g.Group(children),
	)
}