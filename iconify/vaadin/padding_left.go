package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaddingLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 16h16V0H0v16zM3 1v1h1V1h11v14H3v-1H2v1H1v-1h1v-1H1v-1h1v-1H1v-1h1V9H1V8h1V7H1V6h1V5H1V4h1V3H1V2h1V1h1z"/><path fill="currentColor" d="M2 12h1v1H2v-1zm1 1h1v1H3v-1zm0-2h1v1H3v-1zm0-2h1v1H3V9zm-1 1h1v1H2v-1zm0-2h1v1H2V8zm0-2h1v1H2V6zm1 1h1v1H3V7zm0-2h1v1H3V5zm0-2h1v1H3V3zM2 4h1v1H2V4zm0-2h1v1H2V2z"/>`),
		g.Group(children),
	)
}