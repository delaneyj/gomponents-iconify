package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextInput(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2 2h1v4H2V2z"/><path fill="currentColor" d="M1 0C.4 0 0 .4 0 1v14c0 .6.4 1 1 1h15V0H1zm12 15H1V1h12v14zm2 0h-1v-1h1v1zm0-2h-1V3h1v10zm0-11h-1V1h1v1z"/>`),
		g.Group(children),
	)
}