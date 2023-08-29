package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopwatchCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13 19.5a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11Zm0 1a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6.646 7.646a.5.5 0 0 1 .708 0l1.5 1.5a.5.5 0 1 1-.708.708l-1.5-1.5a.5.5 0 0 1 0-.708Zm10.5 2.208a.5.5 0 0 0 .708 0l1.5-1.5a.5.5 0 0 0-.708-.708l-1.5 1.5a.5.5 0 0 0 0 .708Zm-1.256 1.834a.5.5 0 0 1-.078.702l-2.5 2a.5.5 0 0 1-.624-.78l2.5-2a.5.5 0 0 1 .702.078ZM12.5 8V6h1v2h-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M11 6a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1h-3A.5.5 0 0 1 11 6Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}