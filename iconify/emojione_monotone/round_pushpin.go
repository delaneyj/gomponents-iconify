package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundPushpin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M57.031 6.969C50.402.344 39.662.344 33.037 6.971c-5.424 5.426-6.4 13.604-2.943 20.028L2 62l34.998-28.095c6.425 3.458 14.605 2.483 20.033-2.942c6.625-6.627 6.625-17.367 0-23.994"/>`),
		g.Group(children),
	)
}