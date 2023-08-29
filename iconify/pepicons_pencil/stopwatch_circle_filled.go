package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopwatchCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilStopwatchCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M13 19.5a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11Zm0 1a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13Z"/><path d="M6.646 7.646a.5.5 0 0 1 .708 0l1.5 1.5a.5.5 0 1 1-.708.708l-1.5-1.5a.5.5 0 0 1 0-.708Zm10.5 2.208a.5.5 0 0 0 .708 0l1.5-1.5a.5.5 0 0 0-.708-.708l-1.5 1.5a.5.5 0 0 0 0 .708Zm-1.256 1.834a.5.5 0 0 1-.078.702l-2.5 2a.5.5 0 0 1-.624-.78l2.5-2a.5.5 0 0 1 .702.078ZM12.5 8V6h1v2h-1Z"/><path d="M11 6a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1h-3A.5.5 0 0 1 11 6Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilStopwatchCircleFilled0)"/></g>`),
		g.Group(children),
	)
}