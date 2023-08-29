package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkRemoveF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.414 9.858l1.414-1.414a1 1 0 0 0-1.414-1.415L7 8.444L5.586 7.029a1 1 0 1 0-1.414 1.415l1.414 1.414l-1.414 1.414a1 1 0 1 0 1.414 1.414L7 11.272l1.414 1.414a1 1 0 0 0 1.414-1.414L8.414 9.858zM3 .858h8a3 3 0 0 1 3 3v15a2 2 0 0 1-3.348 1.477l-2.978-2.717a1 1 0 0 0-1.348 0l-2.978 2.717A2 2 0 0 1 0 18.858v-15a3 3 0 0 1 3-3z"/>`),
		g.Group(children),
	)
}