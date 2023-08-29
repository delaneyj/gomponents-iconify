package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarginBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 0v14h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1V0H0zm15 12H1V1h14v11zM0 15h1v1H0v-1zm1-1h1v1H1v-1zm1 1h1v1H2v-1zm1-1h1v1H3v-1zm1 1h1v1H4v-1zm1-1h1v1H5v-1zm1 1h1v1H6v-1zm1-1h1v1H7v-1zm1 1h1v1H8v-1zm1-1h1v1H9v-1zm1 1h1v1h-1v-1zm1-1h1v1h-1v-1zm1 1h1v1h-1v-1zm1-1h1v1h-1v-1zm1 1h1v1h-1v-1zm1-1h1v1h-1v-1z"/>`),
		g.Group(children),
	)
}