package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VanishQuarter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 3h2v5h-2V3M4.9 6.3l1.4-1.4l2.8 2.8l-1.3 1.5l-2.9-2.9M8 13H3v-2h5v2"/>`),
		g.Group(children),
	)
}