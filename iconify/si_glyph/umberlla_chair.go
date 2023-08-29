package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmberllaChair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M9 0L1 6v1.984h.953v-1H4v1h1.969v-1h2.047v1h1v5.985h.948V7.984h.005v-1h2.078v1H14v-1h2.016v1H17V6L9 0z"/><path d="M3.229 11.016H1V12h1.678l3.342 2.994l1.011-.01v.985h.938v-.985h1.062v.985H10v-.985h4.047v.985H15v-.985h.984V14H6.562l-3.333-2.984z"/></g>`),
		g.Group(children),
	)
}