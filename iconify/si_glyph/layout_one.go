package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M17 16H1V0h16v16zM2 15h14V1H2v14z"/><path d="M3 2h6v12H3zm7 0h5v6h-5zm0 7h5v5h-5z"/></g>`),
		g.Group(children),
	)
}