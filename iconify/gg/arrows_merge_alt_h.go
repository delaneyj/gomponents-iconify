package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsMergeAltH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.503 6h2v5h4.172L5.846 9.172l1.415-1.415L11.503 12l-4.242 4.243l-1.415-1.415L7.675 13H3.503v5h-2V6Zm18.994 0h2v12h-2v-5h-4.172l1.829 1.829l-1.415 1.414L12.497 12l4.242-4.243l1.415 1.415L16.325 11h4.172V6Z"/>`),
		g.Group(children),
	)
}