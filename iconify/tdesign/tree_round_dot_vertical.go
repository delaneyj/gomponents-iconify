package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeRoundDotVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.126 4A4.002 4.002 0 0 1 23 5a4 4 0 0 1-7.874 1H14a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h1.126A4.002 4.002 0 0 1 23 19a4 4 0 0 1-7.874 1H14a3 3 0 0 1-3-3v-4H8.874A4.002 4.002 0 0 1 1 12a4 4 0 0 1 7.874-1H11V7a3 3 0 0 1 3-3h1.126ZM19 3a2 2 0 1 0 0 4a2 2 0 0 0 0-4ZM5 10a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm14 7a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z"/>`),
		g.Group(children),
	)
}