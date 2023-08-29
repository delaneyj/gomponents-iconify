package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeaveCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M18.346 10.116a.5.5 0 0 1 .705.064l2.083 2.5a.5.5 0 0 1-.768.64l-2.084-2.5a.5.5 0 0 1 .064-.704Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M18.346 15.884a.5.5 0 0 1-.064-.704l2.084-2.5a.5.5 0 1 1 .768.64l-2.083 2.5a.5.5 0 0 1-.704.064Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M20.5 13a.5.5 0 0 1-.5.5h-7.5a.5.5 0 0 1 0-1H20a.5.5 0 0 1 .5.5Zm-14-7a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1H7a.5.5 0 0 1-.5-.5Zm0 14a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1H7a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M16 5.5a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5Zm0 10a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-1 0v-4a.5.5 0 0 1 .5-.5Zm-9-10a.5.5 0 0 1 .5.5v14a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}