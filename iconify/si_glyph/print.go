package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Print(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M6 12v3.98h5.993V12H6zM5 1h7.948v2.96H5z"/><path d="M1.041 5.016v9h3.91V11.01H13v3.006h4.041v-9h-16zm2.975 2.013H2.969V5.953h1.047v1.076zm2-.06H4.969v-1h1.047v1z"/></g>`),
		g.Group(children),
	)
}