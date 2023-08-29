package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.234 4.357A.996.996 0 0 0 2 5v12a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V5.01a1.006 1.006 0 0 0-.364-.781a.996.996 0 0 0-.632-.229H3a.997.997 0 0 0-.766.357ZM4 7.414V17a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V7.414l-7.293 7.293a1 1 0 0 1-1.414 0L4 7.414Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}