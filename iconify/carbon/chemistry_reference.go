package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChemistryReference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 20v2h4.586L2 28.586L3.414 30L10 23.414V28h2v-8H4z"/><path fill="currentColor" d="M20 13.67V4h2V2H10v2h2v9.67L9.58 17h2.477L14 14.33V4h4v10.33l7.61 10.46a2.013 2.013 0 0 1-.44 2.82a2.04 2.04 0 0 1-1.19.39H15v2h8.98a4.015 4.015 0 0 0 3.25-6.38Z"/>`),
		g.Group(children),
	)
}