package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 2.5a.5.5 0 1 0-1 0v11a.5.5 0 0 0 1 0v-11ZM2 3.002a1 1 0 0 1 1.579-.816l7 4.963a1 1 0 0 1 .006 1.628l-7 5.037A1 1 0 0 1 2 13.003V3.002Zm8 4.963L3 3.002v10l7-5.037Z"/>`),
		g.Group(children),
	)
}