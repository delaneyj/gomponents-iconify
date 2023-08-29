package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextHeaderOneTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20.182 5.017A1.001 1.001 0 0 1 21 6.033V18a1 1 0 1 1-2 0V9.611a5.663 5.663 0 0 1-2.183 1.338a1 1 0 0 1-.633-1.897c1.129-.377 2.182-1.333 2.858-3.339a.996.996 0 0 1 .278-.446a1 1 0 0 1 .862-.25ZM3 5a1 1 0 0 1 1 1v5h6V6a1 1 0 1 1 2 0v12a1 1 0 1 1-2 0v-5H4v5a1 1 0 1 1-2 0V6a1 1 0 0 1 1-1Z"/>`),
		g.Group(children),
	)
}