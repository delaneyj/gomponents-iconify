package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentDataSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.5 13a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-1 0v6a.5.5 0 0 0 .5.5Zm5.5-.5a.5.5 0 0 1-1 0v-4a.5.5 0 0 1 1 0v4Zm-3.06.5a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-1 0v2a.5.5 0 0 0 .5.5ZM3 3a2 2 0 0 1 2-2h3.586a1.5 1.5 0 0 1 1.061.439l2.914 2.914c.281.282.439.663.439 1.061V13a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V3Zm2-1a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V6H9.5A1.5 1.5 0 0 1 8 4.5V2H5Zm4.5 3h2.293L9 2.207V4.5a.5.5 0 0 0 .5.5Z"/>`),
		g.Group(children),
	)
}