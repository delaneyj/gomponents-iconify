package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoRedo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.707 5.293a1 1 0 1 0-1.414 1.414l1.414-1.414ZM19 9l.707.707a1 1 0 0 0 0-1.414L19 9Zm-3.707 2.293a1 1 0 0 0 1.414 1.414l-1.414-1.414ZM19 18a1 1 0 1 0 0-2v2ZM15.293 6.707l3 3l1.414-1.414l-3-3l-1.414 1.414Zm3 1.586l-3 3l1.414 1.414l3-3l-1.414-1.414ZM19 8H8v2h11V8ZM3 13a5 5 0 0 0 5 5v-2a3 3 0 0 1-3-3H3Zm5-5a5 5 0 0 0-5 5h2a3 3 0 0 1 3-3V8Zm11 8H8v2h11v-2Z"/>`),
		g.Group(children),
	)
}