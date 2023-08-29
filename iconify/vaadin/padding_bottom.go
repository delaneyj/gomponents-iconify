package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaddingBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 16V0H0v16h16zM1 13h1v-1H1V1h14v12h-1v1h1v1h-1v-1h-1v1h-1v-1h-1v1h-1v-1H9v1H8v-1H7v1H6v-1H5v1H4v-1H3v1H2v-1H1v-1z"/><path fill="currentColor" d="M12 13h1v1h-1v-1zm1-1h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H9v-1zm1 1h1v1h-1v-1zm-2 0h1v1H8v-1zm-2 0h1v1H6v-1zm1-1h1v1H7v-1zm-2 0h1v1H5v-1zm-2 0h1v1H3v-1zm1 1h1v1H4v-1zm-2 0h1v1H2v-1z"/>`),
		g.Group(children),
	)
}