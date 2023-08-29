package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineWest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 19l1.41-1.41L5.83 13H22v-2H5.83l4.59-4.59L9 5l-7 7l7 7z"/>`),
		g.Group(children),
	)
}