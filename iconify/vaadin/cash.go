package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 14H2v-1h13V6h1v8z"/><path fill="currentColor" d="M13 4v7H1V4h12zm1-1H0v9h14V3z"/><path fill="currentColor" d="M3 6H2v3h1v1h4a2.5 2.5 0 1 1 0-5H3v1zm8 0V5H7a2.5 2.5 0 1 1 0 5h4V9h1V6h-1z"/>`),
		g.Group(children),
	)
}