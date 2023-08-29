package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagStraightCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M19.58 6H6.5v8h13.08l-2.299-3.283a1.25 1.25 0 0 1 0-1.434L19.58 6ZM6.5 5a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h13.08a1 1 0 0 0 .819-1.573L18.1 10.143a.25.25 0 0 1 0-.286l2.3-3.284A1 1 0 0 0 19.579 5H6.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6 6a.5.5 0 0 1 .5.5v14a.5.5 0 0 1-1 0v-14A.5.5 0 0 1 6 6Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}