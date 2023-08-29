package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnterCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><path d="M11.25 8a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v13a1 1 0 0 1-1 1h-8a1 1 0 0 1-1-1V8Z" opacity=".2"/><path fill-rule="evenodd" d="M11.547 10.116a.5.5 0 0 1 .704.064l2.083 2.5a.5.5 0 0 1-.768.64l-2.083-2.5a.5.5 0 0 1 .064-.704Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M11.546 15.884a.5.5 0 0 1-.064-.704l2.084-2.5a.5.5 0 1 1 .768.64l-2.083 2.5a.5.5 0 0 1-.704.064Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M14 13a.5.5 0 0 1-.5.5H6a.5.5 0 0 1 0-1h7.5a.5.5 0 0 1 .5.5Zm5 7a.5.5 0 0 1-.5.5H9.3a.5.5 0 0 1 0-1h9.2a.5.5 0 0 1 .5.5Zm0-14a.5.5 0 0 1-.5.5H9.326a.5.5 0 1 1 0-1H18.5a.5.5 0 0 1 .5.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9.25 20.5a.5.5 0 0 1-.5-.5v-4a.5.5 0 0 1 1 0v4a.5.5 0 0 1-.5.5Zm0-10a.5.5 0 0 1-.5-.5V6a.5.5 0 0 1 1 0v4a.5.5 0 0 1-.5.5Zm9.35 10a.5.5 0 0 1-.5-.5V6a.5.5 0 0 1 1 0v14a.5.5 0 0 1-.5.5Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}