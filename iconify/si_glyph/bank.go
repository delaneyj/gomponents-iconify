package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 14v-1h-1.002v-1H14V6.032h-.969V12H11V6.032h-1V12H8V6.032H7V12H5V6.032h-.99L4 12H3v1H2v1H1v1h16v-1h-1zM3.021 6h11.958V5h1V4h-1.01L9.031.094L3.031 4H2v1h1.021v1z"/>`),
		g.Group(children),
	)
}