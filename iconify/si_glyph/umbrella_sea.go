package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaSea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.469 1L1 7v1.984h.953v-1H4v1h1.969v-1h2.047v7.047H5.002v1H3v.907h10.938v-.907H12v-1H8.964V8.984h.005v-1h2.078v1H13v-1h2.016v1H16V7L8.469 1z"/>`),
		g.Group(children),
	)
}