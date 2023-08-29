package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowExportLtrTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.75 3.75a.75.75 0 0 1 .75.75v11a.75.75 0 0 1-1.5 0v-11a.75.75 0 0 1 .75-.75ZM5 10a.75.75 0 0 1 .75-.75h10.19l-2.72-2.72a.75.75 0 0 1 1.06-1.06l3.997 3.996a.775.775 0 0 1 .156.223a.747.747 0 0 1-.156.845L14.28 14.53a.75.75 0 1 1-1.06-1.06l2.72-2.72H5.75A.75.75 0 0 1 5 10Z"/>`),
		g.Group(children),
	)
}