package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaddingRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 0H0v16h16V0zm-3 15v-1h-1v1H1V1h12v1h1V1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h-1z"/><path fill="currentColor" d="M13 3h1v1h-1V3zm-1-1h1v1h-1V2zm0 2h1v1h-1V4zm0 2h1v1h-1V6zm1-1h1v1h-1V5zm0 2h1v1h-1V7zm0 2h1v1h-1V9zm-1-1h1v1h-1V8zm0 2h1v1h-1v-1zm0 2h1v1h-1v-1zm1-1h1v1h-1v-1zm0 2h1v1h-1v-1z"/>`),
		g.Group(children),
	)
}