package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M11 12h4.958v.918H11z"/><path d="M13 10h.918v4.957H13z"/><path d="M11.917 8.958h3.055c.605-1.135.997-2.431.997-3.896a4.019 4.019 0 0 0-4.011-4.031a4.015 4.015 0 0 0-3.911 3.148a4.054 4.054 0 0 0-3.945-3.148c-2.237 0-4.05 1.824-4.05 4.072c0 6.496 8.005 9.838 8.005 9.838s.785-.324 1.86-.974v-3.05h2V8.958z"/></g>`),
		g.Group(children),
	)
}