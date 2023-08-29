package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowExportSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1.5 3a.5.5 0 0 1 .5.5V12a.5.5 0 0 1-1 0V3.5a.5.5 0 0 1 .5-.5Zm9.146.646a.5.5 0 0 1 .708 0l3.5 3.5a.5.5 0 0 1 0 .708l-3.5 3.5a.5.5 0 0 1-.708-.708L13.293 8H4.5a.5.5 0 0 1 0-1h8.793l-2.647-2.646a.5.5 0 0 1 0-.708Z"/>`),
		g.Group(children),
	)
}