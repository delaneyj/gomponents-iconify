package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlidersCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6 7.75a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1-.5-.5Zm11.38 0a.5.5 0 0 1 .5-.5h1.62a.5.5 0 0 1 0 1h-1.62a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M15.75 9.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0 1a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5ZM6 17.75a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1-.5-.5Zm11.38 0a.5.5 0 0 1 .5-.5h1.62a.5.5 0 0 1 0 1h-1.62a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M15.75 19.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0 1a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5ZM6 12.75a.5.5 0 0 1 .5-.5h2.13a.5.5 0 0 1 0 1H6.5a.5.5 0 0 1-.5-.5Zm6.5 0a.5.5 0 0 1 .5-.5h6.5a.5.5 0 0 1 0 1H13a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10.75 14.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0 1a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}