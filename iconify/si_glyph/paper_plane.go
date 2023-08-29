package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m17 1.042l-5.564 13.912l-3.478-3.477l.695 2.086l-1.623 1.395v-3.395l7.954-8.188l-8.937 6.594L1 8.694l16-7.652z"/>`),
		g.Group(children),
	)
}