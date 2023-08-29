package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryChargeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M12.5 11.5H12h.5Zm-1 1v.5v-.5Zm0-10V2v.5Zm1 1h.5h-.5Zm-12 0H0h.5Zm1-1V3v-.5Zm-1 9H1H.5Zm1 1V12v.5Zm5-3l-.224.447A.5.5 0 0 0 7 9.5h-.5Zm0-4l.224-.447A.5.5 0 0 0 6 5.5h.5Zm-5.5 6v-8H0v8h1ZM1.5 3h10V2h-10v1Zm10.5.5v8h1v-8h-1Zm-.5 8.5h-10v1h10v-1Zm.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1ZM11.5 3a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 11.5 2v1ZM1 3.5a.5.5 0 0 1 .5-.5V2A1.5 1.5 0 0 0 0 3.5h1Zm-1 8A1.5 1.5 0 0 0 1.5 13v-1a.5.5 0 0 1-.5-.5H0ZM15 10V5h-1v5h1ZM2.276 7.947l4 2l.448-.894l-4-2l-.448.894ZM7 9.5v-4H6v4h1Zm-.724-3.553l4 2l.448-.894l-4-2l-.448.894Z"/>`),
		g.Group(children),
	)
}