package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Table(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 1v15h16V1H0zm5 14H1v-2h4v2zm0-3H1v-2h4v2zm0-3H1V7h4v2zm0-3H1V4h4v2zm5 9H6v-2h4v2zm0-3H6v-2h4v2zm0-3H6V7h4v2zm0-3H6V4h4v2zm5 9h-4v-2h4v2zm0-3h-4v-2h4v2zm0-3h-4V7h4v2zm0-3h-4V4h4v2z"/>`),
		g.Group(children),
	)
}