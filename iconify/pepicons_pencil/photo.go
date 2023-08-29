package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Photo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M8 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 1a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/><path d="M7.5 10.678L5.59 13.25H7.5l1.91 1H5.59a1 1 0 0 1-.802-1.596l1.909-2.572a1 1 0 0 1 1.606 0l.955 1.286l-.803.596l-.955-1.286Z"/><path fill-rule="evenodd" d="M14.257 13.25L11.25 9.778L8.243 13.25h6.014Zm-2.251-4.127a1 1 0 0 0-1.512 0l-3.007 3.472a1 1 0 0 0 .756 1.655h6.014a1 1 0 0 0 .756-1.655l-3.007-3.472Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M17 2.5H3a.5.5 0 0 0-.5.5v14a.5.5 0 0 0 .5.5h14a.5.5 0 0 0 .5-.5V3a.5.5 0 0 0-.5-.5Zm-14-1A1.5 1.5 0 0 0 1.5 3v14A1.5 1.5 0 0 0 3 18.5h14a1.5 1.5 0 0 0 1.5-1.5V3A1.5 1.5 0 0 0 17 1.5H3Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}