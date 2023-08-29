package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.984 0L2.016.052l.01 10.927l2.005-.01l-.004 1.984h2.02L4.031 16h9.938V6.012H7.984V0z"/><path d="M9 0v4.969h5L9 0zM4.017 14.059l-1.5 1.884l-1.5-1.884H2.03V12h.954v2.059h1.033z"/></g>`),
		g.Group(children),
	)
}