package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatDouble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 2h19v14H5.82L1 19.443V2Zm2 2v11.557L5.18 14H18V4H3Zm20.5.5v18.443L18.68 19.5H7.5v-2h11.82l2.18 1.557V4.5h2Z"/>`),
		g.Group(children),
	)
}