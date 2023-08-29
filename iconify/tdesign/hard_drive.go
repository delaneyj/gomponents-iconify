package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v10h16V4H4Zm16 12H4v4h16v-4Zm-14.002.998h2.004v2.004H5.998v-2.004Zm3 0h2.004v2.004H8.998v-2.004Z"/>`),
		g.Group(children),
	)
}