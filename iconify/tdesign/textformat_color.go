package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextformatColor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.813 17.525l-.394-.919l-6-14L13.16 2h-2.32l-.26.606l-6 14l-.393.92l1.838.787l.394-.92l1.824-4.254h7.515l1.823 4.255l.394.92l1.838-.789ZM9.791 11.14H9.1L12 4.372l2.9 6.767H9.791ZM19 22h1v-2H4v2h15Z"/>`),
		g.Group(children),
	)
}