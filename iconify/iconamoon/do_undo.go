package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoUndo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.707 6.707a1 1 0 0 0-1.414-1.414l1.414 1.414ZM5 9l-.707-.707a1 1 0 0 0 0 1.414L5 9Zm2.293 3.707a1 1 0 1 0 1.414-1.414l-1.414 1.414ZM5 16a1 1 0 0 0 0 2v-2ZM7.293 5.293l-3 3l1.414 1.414l3-3l-1.414-1.414Zm-3 4.414l3 3l1.414-1.414l-3-3l-1.414 1.414ZM5 10h11V8H5v2Zm14 3a3 3 0 0 1-3 3v2a5 5 0 0 0 5-5h-2Zm-3-3a3 3 0 0 1 3 3h2a5 5 0 0 0-5-5v2ZM5 18h11v-2H5v2Z"/>`),
		g.Group(children),
	)
}