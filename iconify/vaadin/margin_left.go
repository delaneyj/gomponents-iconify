package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarginLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2 0v1h1v1H2v1h1v1H2v1h1v1H2v1h1v1H2v1h1v1H2v1h1v1H2v1h1v1H2v1h1v1h13V0H2zm13 15H4V1h11v14zM0 0h1v1H0V0zm1 1h1v1H1V1zM0 2h1v1H0V2zm1 1h1v1H1V3zM0 4h1v1H0V4zm1 1h1v1H1V5zM0 6h1v1H0V6zm1 1h1v1H1V7zM0 8h1v1H0V8zm1 1h1v1H1V9zm-1 1h1v1H0v-1zm1 1h1v1H1v-1zm-1 1h1v1H0v-1zm1 1h1v1H1v-1zm-1 1h1v1H0v-1zm1 1h1v1H1v-1z"/>`),
		g.Group(children),
	)
}