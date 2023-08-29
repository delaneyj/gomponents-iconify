package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignRoadTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.633 3.938L16 2.562l-2.367-1.457l-4.654.418V.021H7v1.573l-3.979.427v2.833L7 4.428v3.063l-4.622-.437L0 8.375l2.378 1.479l4.622.41v5.674h1.979v-5.597l3.979.513V8.055l-3.979-.377V4.357l4.654-.419z"/>`),
		g.Group(children),
	)
}