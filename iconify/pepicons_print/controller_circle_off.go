package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControllerCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><path d="M9.5 9.5h9A4.5 4.5 0 0 1 23 14v2a4.5 4.5 0 0 1-4.5 4.5h-9A4.5 4.5 0 0 1 5 16v-2a4.5 4.5 0 0 1 4.5-4.5Z" opacity=".2"/><path d="M15.25 13a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5Zm2 2.5a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5Z"/><path fill-rule="evenodd" d="M17.5 7.5h-9A4.5 4.5 0 0 0 4 12v2a4.5 4.5 0 0 0 4.5 4.5h9A4.5 4.5 0 0 0 22 14v-2a4.5 4.5 0 0 0-4.5-4.5ZM5 12a3.5 3.5 0 0 1 3.5-3.5h9A3.5 3.5 0 0 1 21 12v2a3.5 3.5 0 0 1-3.5 3.5h-9A3.5 3.5 0 0 1 5 14v-2Z" clip-rule="evenodd"/><path d="M7 14a1 1 0 1 1 0-2h4a1 1 0 0 1 0 2H7Z"/><path d="M10 15a1 1 0 1 1-2 0v-4a1 1 0 0 1 2 0v4Z"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}