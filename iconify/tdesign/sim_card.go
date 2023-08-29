package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.586 1H21v22H3V5.586L7.586 1Zm.828 2L5 6.414V21h14V3H8.414ZM8 8.998h2.004v2.004H8V8.998Zm6 0h2.004v2.004H14V8.998ZM13 9v5h-2V9h2Zm-3 3v5H8v-5h2Zm6 0v5h-2v-5h2Zm-5.002 2.998h2.004v2.004h-2.004v-2.004Z"/>`),
		g.Group(children),
	)
}