package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MsWordOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m3.5 9.5l-.485.121a.5.5 0 0 0 .901.156L3.5 9.5Zm1-1.5l.416-.277a.5.5 0 0 0-.832 0L4.5 8Zm1 1.5l-.416.277a.5.5 0 0 0 .901-.156L5.5 9.5ZM1.5 4h6V3h-6v1Zm6.5.5v6h1v-6H8ZM7.5 11h-6v1h6v-1ZM1 10.5v-6H0v6h1Zm.5.5a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 12v-1Zm6.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 10.5H8ZM7.5 4a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 3v1Zm-6-1A1.5 1.5 0 0 0 0 4.5h1a.5.5 0 0 1 .5-.5V3Zm.515 2.621l1 4l.97-.242l-1-4l-.97.242Zm1.901 4.156l1-1.5l-.832-.554l-1 1.5l.832.554Zm.168-1.5l1 1.5l.832-.554l-1-1.5l-.832.554Zm1.901 1.344l1-4l-.97-.242l-1 4l.97.242ZM3 3.5v-2H2v2h1ZM3.5 1h10V0h-10v1Zm10.5.5v12h1v-12h-1ZM13.5 14h-10v1h10v-1ZM3 13.5v-2H2v2h1Zm.5.5a.5.5 0 0 1-.5-.5H2A1.5 1.5 0 0 0 3.5 15v-1Zm10.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1ZM13.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 0v1ZM3 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 2 1.5h1Z"/>`),
		g.Group(children),
	)
}