package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M19.5 7.5h-13A.5.5 0 0 0 6 8v11a.5.5 0 0 0 .5.5h13a.5.5 0 0 0 .5-.5V8a.5.5 0 0 0-.5-.5ZM7 18.5v-10h12v10H7Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M19.5 7.5h-13A.5.5 0 0 0 6 8v3a.5.5 0 0 0 .5.5h13a.5.5 0 0 0 .5-.5V8a.5.5 0 0 0-.5-.5ZM7 10.5v-2h12v2H7Z" clip-rule="evenodd"/><path d="M8.5 8.5h1A.5.5 0 0 0 10 8V7a.5.5 0 0 0-.5-.5h-1A.5.5 0 0 0 8 7v1a.5.5 0 0 0 .5.5Zm8 0h1A.5.5 0 0 0 18 8V7a.5.5 0 0 0-.5-.5h-1a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5Zm-7.5 6a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}