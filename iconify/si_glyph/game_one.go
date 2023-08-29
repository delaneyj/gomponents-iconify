package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GameOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m8.931 7.986l4.813-6.343A7.875 7.875 0 0 0 8.931-.001C4.551-.001 1 3.561 1 7.956c0 4.396 3.551 7.959 7.931 7.959c3.517 0 6.495-2.3 7.533-5.481L8.931 7.986zM7.499 5a1.5 1.5 0 1 1 .003-3.001A1.5 1.5 0 0 1 7.499 5z"/>`),
		g.Group(children),
	)
}