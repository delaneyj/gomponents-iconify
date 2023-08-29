package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackspaceFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.92 4H19a3 3 0 0 1 3 3v10a3 3 0 0 1-3 3H7.92a3 3 0 0 1-2.278-1.048l-4.285-5a3 3 0 0 1 0-3.904l4.285-5A3 3 0 0 1 7.92 4Zm2.373 5.293a1 1 0 0 1 1.414 0L13 10.586l1.293-1.293a1 1 0 1 1 1.414 1.414L14.414 12l1.293 1.293a1 1 0 0 1-1.414 1.414L13 13.414l-1.293 1.293a1 1 0 1 1-1.414-1.414L11.586 12l-1.293-1.293a1 1 0 0 1 0-1.414Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}