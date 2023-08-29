package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowResizeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m5.075 6.95l1.843-1.862l-2.98-3.026L5.955.018H.052v5.986l2.046-2.076L5.075 6.95zm10.928 2.966l-2.171 2.097l-3.033-3.053l-1.881 1.881l3.039 3.056l-1.996 2.084h6.042V9.916z"/>`),
		g.Group(children),
	)
}