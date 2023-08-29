package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberFiveFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.01 3.859A1 1 0 0 1 9 3h7a1 1 0 1 1 0 2H9.867L9.26 9.257A6 6 0 1 1 7.4 19.8a1 1 0 1 1 1.202-1.6a4 4 0 1 0 0-6.402a1 1 0 0 1-1.59-.94l.999-7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}