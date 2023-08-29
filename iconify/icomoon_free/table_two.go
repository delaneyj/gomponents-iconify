package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 1v14h16V1H0zm6 9V7h4v3H6zm4 1v3H6v-3h4zm0-8v3H6V3h4zM5 3v3H1V3h4zM1 7h4v3H1V7zm10 0h4v3h-4V7zm0-1V3h4v3h-4zM1 11h4v3H1v-3zm10 3v-3h4v3h-4z"/>`),
		g.Group(children),
	)
}