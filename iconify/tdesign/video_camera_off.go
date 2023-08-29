package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCameraOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1.586L22.414 21L21 22.414l-4-4V20H0V4h2.586l-1-1L3 1.586ZM4.586 6H2v12h13v-1.586L4.586 6Zm2.996-2H17v4.433l7-4.2V16l.003 3.416L22 17.415V7.766l-5 3l.004 2.655L15 11.415V6.001l-5.413.001L7.582 4Z"/>`),
		g.Group(children),
	)
}