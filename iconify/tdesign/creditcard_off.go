package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditcardOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 .586L23.414 22L22 23.414L19.586 21H1V3h.586l-1-1L2 .586ZM3 5v4h4.586l-4-4H3Zm0 6v8h14.586l-8-8H3Zm4.581-8.001L23.001 3l.002 9.418l-.003-.004l.003 6.004L21 16.415v-5.414l-5.413.001l-2.005-2.003L21 9V5l-11.413.001l-2.006-2.003ZM5.001 14h5v2H5v-2Z"/>`),
		g.Group(children),
	)
}