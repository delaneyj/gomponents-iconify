package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineSyncAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.41 13.41L6 12l-4 4l4 4l1.41-1.41L5.83 17H21v-2H5.83zm9.18-2.82L18 12l4-4l-4-4l-1.41 1.41L18.17 7H3v2h15.17z"/>`),
		g.Group(children),
	)
}