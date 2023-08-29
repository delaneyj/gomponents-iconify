package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenQuill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m23.302 8.118l-11.14 11.094l-5.416-.697l-3.673 3.673l-1.414-1.414l3.669-3.67l-.742-5.41L15.672.565l1.816 5.787l5.814 1.767Zm-7.745-1.242l-.803-2.557l-8.052 8.085l.401 2.926l8.454-8.454Zm-7.025 9.853l2.914.375l8.076-8.045l-2.546-.773l-8.444 8.443Z"/>`),
		g.Group(children),
	)
}