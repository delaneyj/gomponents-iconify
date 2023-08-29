package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FitnessCentre(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M14.5 7v1h-1v2h-1v1H11V8H4v3H2.5v-1h-1V8h-1V7h1V5h1V4H4v3h7V4h1.5v1h1v2Z"/>`),
		g.Group(children),
	)
}