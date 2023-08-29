package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Table(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 3v11h16V3H0zm6 7V8h4v2H6zm4 1v2H6v-2h4zm0-6v2H6V5h4zM5 5v2H1V5h4zM1 8h4v2H1V8zm10 0h4v2h-4V8zm0-1V5h4v2h-4zM1 11h4v2H1v-2zm10 2v-2h4v2h-4z"/>`),
		g.Group(children),
	)
}