package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Map(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M373 0q11 0 11 11v322q0 8-8 10l-120 41l-128-45l-114 44l-3 1q-11 0-11-11V51q0-8 8-10L128 0l128 45L370 1zM256 341V88L128 43v253z"/>`),
		g.Group(children),
	)
}