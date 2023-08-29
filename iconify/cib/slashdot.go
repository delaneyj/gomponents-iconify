package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slashdot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26.37 0L9.38 32H2.489L19.473 0zm3.14 26.401a5.478 5.478 0 0 1-5.479 5.474c-3.026 0-5.479-2.453-5.479-5.474s2.453-5.474 5.479-5.474a5.478 5.478 0 0 1 5.479 5.474z"/>`),
		g.Group(children),
	)
}