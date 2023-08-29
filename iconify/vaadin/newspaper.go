package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Newspaper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2 4h11v4H2V4zm0-2h11v1H2V2zm6 11h3v1H8v-1zm0-2h5v1H8v-1zm0-2h5v1H8V9zm-6 4h5v1H2v-1zm0-2h5v1H2v-1zm0-2h5v1H2V9z"/><path fill="currentColor" d="M15 2V0H0v14.5A1.5 1.5 0 0 0 1.5 16h13a1.5 1.5 0 0 0 1.5-1.5V2h-1zM1.5 15a.5.5 0 0 1-.5-.5V1h13v12.5c0 1.5 1 1.5 1 1.5H1.5z"/>`),
		g.Group(children),
	)
}