package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberFourFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.196 3.02a1 1 0 0 1 .785 1.176l-1.522 7.608a1 1 0 0 0 .98 1.196H15V8a1 1 0 1 1 2 0v12a1 1 0 1 1-2 0v-5h-4.56a3 3 0 0 1-2.942-3.588l1.521-7.608a1 1 0 0 1 1.177-.785Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}