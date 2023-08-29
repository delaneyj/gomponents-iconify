package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowResizeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.995 10.852L5.133 9.008l-3.026 2.98L.062 9.972v5.903h5.987l-2.076-2.047l3.022-2.976zM9.961.008l2.097 2.087l-3.053 3.033l1.88 1.88l3.057-3.038l1.967 1.996V.008H9.961z"/>`),
		g.Group(children),
	)
}