package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextFootnote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 7v2h7v16h2V9h7V7H2zm28 4.076l-.744-1.857L26 10.522V7h-2v3.523L20.744 9.22L20 11.077l3.417 1.367L20.9 15.8l1.6 1.2l2.5-3.333L27.5 17l1.6-1.2l-2.517-3.357L30 11.076z"/>`),
		g.Group(children),
	)
}