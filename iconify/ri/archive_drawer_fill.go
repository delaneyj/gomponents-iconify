package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveDrawerFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 13h18v8.003c0 .55-.445.997-.993.997H3.993A.995.995 0 0 1 3 21.003V13ZM3 2.998C3 2.447 3.445 2 3.993 2h16.014c.548 0 .993.446.993.998V11H3V2.998ZM9 5v2h6V5H9Zm0 11v2h6v-2H9Z"/>`),
		g.Group(children),
	)
}