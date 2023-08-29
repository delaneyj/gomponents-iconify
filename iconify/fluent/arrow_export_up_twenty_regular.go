package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowExportUpTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.646 2.146a.5.5 0 0 1 .708 0l4 4a.5.5 0 0 1-.708.708L10.5 3.707V14.5a.5.5 0 0 1-1 0V3.707L6.354 6.854a.5.5 0 1 1-.708-.708l4-4ZM4 17.5a.5.5 0 0 1 .5-.5h11a.5.5 0 0 1 0 1h-11a.5.5 0 0 1-.5-.5Z"/>`),
		g.Group(children),
	)
}