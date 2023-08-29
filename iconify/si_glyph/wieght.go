package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wieght(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 6.079V5h-1.955v3H4V5.041H2v1.011h-.961v5.906H2v.997h2V10h10.045v2.996H16v-1.038h1V6.079h-1z"/>`),
		g.Group(children),
	)
}