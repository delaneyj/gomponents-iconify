package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Install(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 2v7.11l2.508-2.48l1.406 1.422L12 12.91L7.086 8.052L8.492 6.63L11 9.11V2h2ZM2 2h7v2H4v10h16V4h-5V2h7v20H2V2Zm18 14H4v4h16v-4Zm-14.002.998h2.004v2.004H5.998v-2.004Zm3 0h2.004v2.004H8.998v-2.004Z"/>`),
		g.Group(children),
	)
}