package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StreetTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17.02 14.046v-13h-16v13h16zM14 6.953h2v1h-2v-1zM9 1.992h4V4.02H9V1.992zM9 5h4.021v2H9V5zm0 3h4.021v2H9V8zm0 2.958h4v2H9v-2zM5 6.953h2V8H5V6.953zm-3 0h2V8H2V6.953z"/>`),
		g.Group(children),
	)
}