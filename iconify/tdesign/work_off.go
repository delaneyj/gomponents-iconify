package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorkOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 .586L23.414 22L22 23.414L20.086 21.5H2v-15h3.086L.586 2L2 .586ZM7.086 8.5H4v11h14.086l-11-11Zm.42-6H16.5v4.002L22 6.5l.003 11.418L20 15.915V8.501h-7.413l-2.003-2.003l3.916.004V4.5h-5v1.622L7.497 4.078L7.506 2.5Z"/>`),
		g.Group(children),
	)
}