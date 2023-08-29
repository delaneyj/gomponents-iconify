package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Balance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M43 197h64v150H43V197zm128 0h64v150h-64V197zM0 453v-64h405v64H0zm299-256h64v150h-64V197zM203 5l202 107v43H0v-43z"/>`),
		g.Group(children),
	)
}