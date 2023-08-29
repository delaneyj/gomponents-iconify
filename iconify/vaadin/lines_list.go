package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinesList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 1h3v2H0V1zm0 4h3v2H0V5zm0 4h3v2H0V9zm0 4h3v2H0v-2zM4 1h12v2H4V1zm0 4h12v2H4V5zm0 4h12v2H4V9zm0 4h12v2H4v-2z"/>`),
		g.Group(children),
	)
}