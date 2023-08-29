package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MeasurementOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h10v10h10v10H2V2Zm2 2v2.5h2.004v2H4V11h2.004v2H4v2.5h2.004v2H4V20h2.5v-1.967h2V20H11v-1.967h2V20h2.5v-1.967h2V20H20v-6H10V4H4Z"/>`),
		g.Group(children),
	)
}