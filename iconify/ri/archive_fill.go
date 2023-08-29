package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 10h18v10.004c0 .55-.445.996-.993.996H3.993A.994.994 0 0 1 3 20.004V10Zm6 2v2h6v-2H9ZM2 4c0-.552.455-1 .992-1h18.016c.548 0 .992.444.992 1v4H2V4Z"/>`),
		g.Group(children),
	)
}