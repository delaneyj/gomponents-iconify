package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartColumn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M16 14.031H.984V1.016H.027L0 14.95h.027v.029L16 14.95v-.919z"/><path d="M4.958 8.021H2.016v4.964h2.942V8.021zm5.011-1.974H7.016v6.922h2.953V6.047zm4.984-2.016H12v8.947h2.953V4.031z"/></g>`),
		g.Group(children),
	)
}