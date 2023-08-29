package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.978 1.162c0 .225-.062.45-.196.65L6.626 8.041l4.197 6.037c.359.541.213 1.27-.328 1.629a1.174 1.174 0 0 1-1.63-.325l-4.63-6.688a1.172 1.172 0 0 1-.002-1.304L8.822.51a1.178 1.178 0 0 1 2.156.652z"/>`),
		g.Group(children),
	)
}