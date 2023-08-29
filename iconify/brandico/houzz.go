package brandico

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Houzz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M288.672 0v333.329L0 500.012V166.635L288.672 0M.006 499.988v333.323L288.678 666.58L.036 499.987h-.03M288.672 666.58v333.377l288.666-166.671V499.987L288.672 666.58m0-333.227l288.672 166.635V166.635L288.672 333.354"/>`),
		g.Group(children),
	)
}