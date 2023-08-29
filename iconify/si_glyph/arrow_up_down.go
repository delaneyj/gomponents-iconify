package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.021 3.961V12H5.062l3.819 3.991L12.896 12H10V3.961h2.896L8.857.105L4.991 3.961h3.03z"/>`),
		g.Group(children),
	)
}