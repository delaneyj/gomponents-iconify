package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MpThreeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M13.5 3.5h.5v-.207l-.146-.147l-.354.354Zm-3-3l.354-.354L10.707 0H10.5v.5Zm-8 6l.354-.354A.5.5 0 0 0 2 6.5h.5Zm1 1l-.354.354l.354.353l.354-.353L3.5 7.5Zm1-1H5a.5.5 0 0 0-.854-.354L4.5 6.5Zm2 0V6H6v.5h.5Zm6 0l.4.3a.5.5 0 0 0-.4-.8v.5Zm-1.5 2l-.4-.3a.5.5 0 0 0 .4.8v-.5ZM2 5V1.5H1V5h1Zm11-1.5V5h1V3.5h-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1ZM1 12v1.5h1V12H1Zm1.5 3h10v-1h-10v1ZM14 13.5V12h-1v1.5h1ZM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5H1ZM3 11V6.5H2V11h1Zm-.854-4.146l1 1l.708-.708l-1-1l-.708.708Zm1.708 1l1-1l-.708-.708l-1 1l.708.708ZM4 6.5V11h1V6.5H4Zm2.5.5h1V6h-1v1Zm.5 4V8.5H6V11h1Zm0-2.5v-2H6v2h1Zm.5-.5h-1v1h1V8Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 7.5H8ZM7.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 6v1ZM10 7h2.5V6H10v1Zm2.1-.8l-1.5 2l.8.6l1.5-2l-.8-.6ZM11 9h.5V8H11v1Zm.5 1H10v1h1.5v-1Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 13 9.5h-1Zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 11.5 8v1Z"/>`),
		g.Group(children),
	)
}