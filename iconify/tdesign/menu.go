package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Menu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v5.5h16V4H4Zm16 7.5H4V20h16v-8.5ZM5.996 6H8v2h-.004v.004h-2V6ZM10 6h8v2h-8V6Z"/>`),
		g.Group(children),
	)
}