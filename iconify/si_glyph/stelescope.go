package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stelescope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="m2.606 3.764l-2.2-2.2L1.8.17L4 2.372z"/><path d="M1.082 3.647L4.33 6.896l2.477-2.477L3.559 1.17L1.082 3.647zm3.623 4.532l4.78 5.729l4.247-4.248l-5.729-4.78l-3.298 3.299zm11.607 2.901l-5.15 5.15l-.81-.808l5.15-5.151z"/></g>`),
		g.Group(children),
	)
}