package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22h14c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2h-2V2h-2v2H9V2H7v2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2zm10.707-4.707l-1.414 1.414L12 16.414l-2.293 2.293l-1.414-1.414L10.586 15l-2.293-2.293l1.414-1.414L12 13.586l2.293-2.293l1.414 1.414L13.414 15l2.293 2.293zM5 7h14v2H5V7z"/>`),
		g.Group(children),
	)
}