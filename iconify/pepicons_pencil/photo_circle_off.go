package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M11 11a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 1a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/><path d="M10.5 13.678L8.59 16.25h1.91l1.91 1H8.59a1 1 0 0 1-.802-1.596l1.909-2.572a1 1 0 0 1 1.606 0l.955 1.286l-.803.596l-.955-1.286Z"/><path fill-rule="evenodd" d="m17.257 16.25l-3.007-3.472l-3.007 3.472h6.014Zm-2.251-4.127a1 1 0 0 0-1.512 0l-3.007 3.472a1 1 0 0 0 .756 1.655h6.014a1 1 0 0 0 .756-1.655l-3.007-3.472Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M20 5.5H6a.5.5 0 0 0-.5.5v14a.5.5 0 0 0 .5.5h14a.5.5 0 0 0 .5-.5V6a.5.5 0 0 0-.5-.5Zm-14-1A1.5 1.5 0 0 0 4.5 6v14A1.5 1.5 0 0 0 6 21.5h14a1.5 1.5 0 0 0 1.5-1.5V6A1.5 1.5 0 0 0 20 4.5H6Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}