package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Widget(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.35 3h13.3L23 12.788V22H1v-9.212L5.35 3Zm1.3 2l-3.111 7H20.46l-3.11-7H6.65ZM21 14H3v6h18v-6ZM5 16h2.004v2.004H5V16Zm4 0h2.004v2.004H9V16Z"/>`),
		g.Group(children),
	)
}