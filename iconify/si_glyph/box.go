package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Box(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 2h16v2H1zm1 10a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V5H2v7zm4.98-5.041h4.082v1.104H6.98V6.959z"/>`),
		g.Group(children),
	)
}