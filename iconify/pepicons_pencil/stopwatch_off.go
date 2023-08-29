package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopwatchOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M10 16.5a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11Zm0 1a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M3.646 4.646a.5.5 0 0 1 .708 0l1.5 1.5a.5.5 0 1 1-.708.708l-1.5-1.5a.5.5 0 0 1 0-.708Zm10.5 2.208a.5.5 0 0 0 .708 0l1.5-1.5a.5.5 0 0 0-.708-.708l-1.5 1.5a.5.5 0 0 0 0 .708ZM12.89 8.688a.5.5 0 0 1-.078.702l-2.5 2a.5.5 0 0 1-.624-.78l2.5-2a.5.5 0 0 1 .702.078ZM9.5 5V3h1v2h-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8 3a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1h-3A.5.5 0 0 1 8 3Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}