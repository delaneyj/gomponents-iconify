package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 2v12h16V2H0zm15 11H1V8h14v5zm0-8H1V3h14v2z"/><path fill="currentColor" d="M10 11h3v1h-3v-1zm-8 0h6v1H2v-1z"/>`),
		g.Group(children),
	)
}